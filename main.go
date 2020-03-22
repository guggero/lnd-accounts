package main

import (
	"fmt"
	"github.com/lightninglabs/aperture/plugin/shared"
	"path"

	"github.com/guggero/lnd-accounts/acctrpc"
	"github.com/guggero/lnd-accounts/macaroons"
	"github.com/lightninglabs/aperture/plugin/author"
	"github.com/lightninglabs/aperture/plugin/pluginrpc"
	"github.com/lightninglabs/loop/lndclient"
	"github.com/lightningnetwork/lnd/lncfg"
	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/grpc"
)

const (
	PluginName = "lnd-accounts"
	DataDir    = "~/.aperture/plugins/accounts"
)

func main() {
	p := &LndAccountsPlugin{}
	author.RunPluginServer(PluginName, p, &shared.RpcPackage{
		Name:     "acctrpc",
		Services: []interface{}{
			(*acctrpc.AccountsServer)(nil),
		},
	})
}

type LndAccountsPlugin struct {
	*RpcServer

	lnd lnrpc.LightningClient
}

func (a *LndAccountsPlugin) Init(request *pluginrpc.StartRequest) error {
	macService, err := macaroons.NewService(lncfg.CleanAndExpandPath(DataDir))
	if err != nil {
		return err
	}

	a.RpcServer = &RpcServer{
		macService: macService,
	}
	a.lnd, err = lndclient.NewBasicClient(
		request.LndHost, request.LndCert, path.Dir(request.LndMacaroon),
		"", lndclient.MacFilename(path.Base(request.LndMacaroon)),
	)
	log.Debugf("Successfully connected to lnd at %s", request.LndHost)
	return err
}

func (a *LndAccountsPlugin) RegisterServices(s *grpc.Server) {
	acctrpc.RegisterAccountsServer(s, a)
}

func (a *LndAccountsPlugin) Stop() error {
	return nil
}

func (a *LndAccountsPlugin) HookLightningNewAddressReq(req *lnrpc.NewAddressRequest) error {
	fmt.Printf("Intercepted request with type: %v\n", req)
	req.Type = lnrpc.AddressType_NESTED_PUBKEY_HASH
	return nil
}

func (a *LndAccountsPlugin) HookLightningNewAddressRes(res *lnrpc.NewAddressResponse) error {
	fmt.Printf("Intercepted response with new address %s\n", res.Address)
	res.Address = "fooooo"
	return nil
}

func (a *LndAccountsPlugin) HookLightningOpenChannelReq(req *lnrpc.OpenChannelRequest) error {
	fmt.Printf("Intercepted request with type: %v\n", req)
	return nil
}

func (a *LndAccountsPlugin) HookLightningOpenChannelRes(res *lnrpc.OpenStatusUpdate) error {
	fmt.Printf("Intercepted response with type: %v\n", res)
	return nil
}

var _ author.PluginImplementer = (*LndAccountsPlugin)(nil)
