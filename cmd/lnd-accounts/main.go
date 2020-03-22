// Copyright (c) 2013-2017 The btcsuite developers
// Copyright (c) 2015-2016 The Decred developers
// Copyright (C) 2015-2020 The Lightning Network Developers

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/btcsuite/btcutil"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/guggero/lnd-accounts/acctrpc"
	"github.com/lightningnetwork/lnd/lncfg"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lightningnetwork/lnd/macaroons"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gopkg.in/macaroon.v2"
)

const (
	defaultDataDir          = "data"
	defaultChainSubDir      = "chain"
	defaultTLSCertFilename  = "tls.cert"
	defaultMacaroonFilename = "admin.macaroon"
	defaultRPCPort          = "11010"
	defaultRPCHostPort      = "localhost:" + defaultRPCPort
)

var (
	defaultLndDir      = btcutil.AppDataDir("lnd", false)
	defaultTLSCertPath = filepath.Join(defaultLndDir, defaultTLSCertFilename)

	// maxMsgRecvSize is the largest message our client will receive. We
	// set this to 200MiB atm.
	maxMsgRecvSize = grpc.MaxCallRecvMsgSize(1 * 1024 * 1024 * 200)
)

func fatal(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "[lnd-accounts] %v\n", err)
	os.Exit(1)
}

func getClient(ctx *cli.Context) (lnrpc.LightningClient, acctrpc.AccountsClient,
	func()) {

	conn := getClientConn(ctx)

	cleanUp := func() {
		err := conn.Close()
		if err != nil {
			fatal(err)
		}
	}

	return lnrpc.NewLightningClient(conn), acctrpc.NewAccountsClient(conn),
		cleanUp
}

func getClientConn(ctx *cli.Context) *grpc.ClientConn {
	// First, we'll parse the args from the command.
	tlsCertPath, macPath, err := extractPathArgs(ctx)
	if err != nil {
		fatal(err)
	}

	// Load the specified TLS certificate and build transport credentials
	// with it.
	creds, err := credentials.NewClientTLSFromFile(tlsCertPath, "")
	if err != nil {
		fatal(err)
	}

	// Create a dial options array.
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	// Load the specified macaroon file.
	macBytes, err := ioutil.ReadFile(macPath)
	if err != nil {
		fatal(fmt.Errorf("unable to read macaroon path (check "+
			"the network setting!): %v", err))
	}

	mac := &macaroon.Macaroon{}
	if err = mac.UnmarshalBinary(macBytes); err != nil {
		fatal(fmt.Errorf("unable to decode macaroon: %v", err))
	}

	macConstraints := []macaroons.Constraint{
		// We add a time-based constraint to prevent replay of the
		// macaroon. It's good for 60 seconds by default to make up for
		// any discrepancy between client and server clocks, but leaking
		// the macaroon before it becomes invalid makes it possible for
		// an attacker to reuse the macaroon. In addition, the validity
		// time of the macaroon is extended by the time the server clock
		// is behind the client clock, or shortened by the time the
		// server clock is ahead of the client clock (or invalid
		// altogether if, in the latter case, this time is more than 60
		// seconds).
		// TODO(aakselrod): add better anti-replay protection.
		macaroons.TimeoutConstraint(ctx.GlobalInt64("macaroontimeout")),

		// Lock macaroon down to a specific IP address.
		macaroons.IPLockConstraint(ctx.GlobalString("macaroonip")),

		// ... Add more constraints if needed.
	}

	// Apply constraints to the macaroon.
	constrainedMac, err := macaroons.AddConstraints(mac, macConstraints...)
	if err != nil {
		fatal(err)
	}

	// Now we append the macaroon credentials to the dial options.
	cred := macaroons.NewMacaroonCredential(constrainedMac)
	opts = append(opts, grpc.WithPerRPCCredentials(cred))

	// We need to use a custom dialer so we can also connect to unix sockets
	// and not just TCP addresses.
	genericDialer := lncfg.ClientAddressDialer(defaultRPCPort)
	opts = append(opts, grpc.WithContextDialer(genericDialer))
	opts = append(opts, grpc.WithDefaultCallOptions(maxMsgRecvSize))

	conn, err := grpc.Dial(ctx.GlobalString("rpcserver"), opts...)
	if err != nil {
		fatal(fmt.Errorf("unable to connect to RPC server: %v", err))
	}

	return conn
}

// extractPathArgs parses the TLS certificate and macaroon paths from the
// command.
func extractPathArgs(ctx *cli.Context) (string, string, error) {
	// We'll start off by parsing the active chain and network. These are
	// needed to determine the correct path to the macaroon when not
	// specified.
	chain := strings.ToLower(ctx.GlobalString("chain"))
	switch chain {
	case "bitcoin", "litecoin":
	default:
		return "", "", fmt.Errorf("unknown chain: %v", chain)
	}

	network := strings.ToLower(ctx.GlobalString("network"))
	switch network {
	case "mainnet", "testnet", "regtest", "simnet":
	default:
		return "", "", fmt.Errorf("unknown network: %v", network)
	}

	// We'll now fetch the lnddir so we can make a decision  on how to
	// properly read the macaroons (if needed) and also the cert. This will
	// either be the default, or will have been overwritten by the end
	// user.
	lndDir := lncfg.CleanAndExpandPath(ctx.GlobalString("lnddir"))

	// If the macaroon path as been manually provided, then we'll only
	// target the specified file.
	var macPath string
	if ctx.GlobalString("macaroonpath") != "" {
		macPath = lncfg.CleanAndExpandPath(
			ctx.GlobalString("macaroonpath"),
		)
	} else {
		// Otherwise, we'll go into the path:
		// lnddir/data/chain/<chain>/<network> in order to fetch the
		// macaroon that we need.
		macPath = filepath.Join(
			lndDir, defaultDataDir, defaultChainSubDir, chain,
			network, defaultMacaroonFilename,
		)
	}

	tlsCertPath := lncfg.CleanAndExpandPath(
		ctx.GlobalString("tlscertpath"),
	)

	// If a custom lnd directory was set, we'll also check if custom paths
	// for the TLS cert and macaroon file were set as well. If not, we'll
	// override their paths so they can be found within the custom lnd
	// directory set. This allows us to set a custom lnd directory, along
	// with custom paths to the TLS cert and macaroon file.
	if lndDir != defaultLndDir {
		tlsCertPath = filepath.Join(lndDir, defaultTLSCertFilename)
	}

	return tlsCertPath, macPath, nil
}

func printRespJSON(resp proto.Message) {
	jsonMarshaler := &jsonpb.Marshaler{
		EmitDefaults: true,
		Indent:       "    ",
	}

	jsonStr, err := jsonMarshaler.MarshalToString(resp)
	if err != nil {
		fmt.Println("unable to decode response: ", err)
		return
	}

	fmt.Println(jsonStr)
}

func main() {
	app := cli.NewApp()
	app.Name = "lnd-accounts"
	app.Version = "0.0.1"
	app.Usage = "control plane for lnd accounts"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "rpcserver",
			Value: defaultRPCHostPort,
			Usage: "host:port of ln daemon",
		},
		cli.StringFlag{
			Name:  "lnddir",
			Value: defaultLndDir,
			Usage: "path to lnd's base directory",
		},
		cli.StringFlag{
			Name:  "tlscertpath",
			Value: defaultTLSCertPath,
			Usage: "path to TLS certificate",
		},
		cli.StringFlag{
			Name:  "chain, c",
			Usage: "the chain lnd is running on e.g. bitcoin",
			Value: "bitcoin",
		},
		cli.StringFlag{
			Name: "network, n",
			Usage: "the network lnd is running on e.g. mainnet, " +
				"testnet, etc.",
			Value: "mainnet",
		},
		cli.StringFlag{
			Name:  "macaroonpath",
			Usage: "path to macaroon file",
		},
		cli.Int64Flag{
			Name:  "macaroontimeout",
			Value: 60,
			Usage: "anti-replay macaroon validity time in seconds",
		},
		cli.StringFlag{
			Name:  "macaroonip",
			Usage: "if set, lock macaroon to specific IP address",
		},
	}
	app.Commands = []cli.Command{
		accountMacaroonCommand,
		createAccountCommand,
		listAccountsCommand,
		removeAccountCommand,
	}

	if err := app.Run(os.Args); err != nil {
		fatal(err)
	}
}
