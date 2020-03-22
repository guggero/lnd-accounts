package macaroons

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"path"

	"github.com/coreos/bbolt"
	"github.com/lightningnetwork/lnd/lntypes"
	"github.com/lightningnetwork/lnd/lnwire"
	"google.golang.org/grpc/metadata"
	"gopkg.in/macaroon.v2"
)

var (
	// DBFilename is the filename within the data directory which contains
	// the macaroon stores.
	DBFilename = "accounts.db"
)

// Service encapsulates bakery.Bakery and adds a Close() method that zeroes the
// root key service encryption keys, as well as utility methods to validate a
// macaroon against the bakery and gRPC middleware for macaroon-based auth.
// Additionally, there is an account storage for accounting based macaroon
// balances and utility methods to manage accounts.
type Service struct {
	*AccountStorage
}

// NewService returns a service backed by the macaroon Bolt DB stored in the
// passed directory.
func NewService(dir string) (*Service, error) {
	// Ensure that the path to the directory exists.
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0700); err != nil {
			return nil, err
		}
	}

	// Open the database that we'll use to store the primary macaroon key,
	// and all generated macaroons+caveats.
	accountsDB, err := bbolt.Open(
		path.Join(dir, DBFilename), 0600, bbolt.DefaultOptions,
	)
	if err != nil {
		return nil, err
	}

	accountStore, err := NewAccountStorage(accountsDB)
	if err != nil {
		return nil, err
	}
	return &Service{accountStore}, nil
}

// macaroonFromContext reads the macaroon content from a context's metadata and
// unmarshals it.
func macaroonFromContext(ctx context.Context) (*macaroon.Macaroon, error) {
	// Get macaroon bytes from context and unmarshal into macaroon.
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("unable to get metadata from context")
	}
	if len(md["macaroon"]) != 1 {
		return nil, fmt.Errorf(
			"expected 1 macaroon, got %d", len(md["macaroon"]),
		)
	}

	// With the macaroon obtained, we'll now decode the hex-string
	// encoding, then unmarshal it from binary into its concrete struct
	// representation.
	macBytes, err := hex.DecodeString(md["macaroon"][0])
	if err != nil {
		return nil, err
	}
	mac := &macaroon.Macaroon{}
	err = mac.UnmarshalBinary(macBytes)
	if err != nil {
		return nil, err
	}
	return mac, nil
}

// ValidateAccountBalance reads the macaroon from the context and checks
// if it is locked to an account. If it is locked, the account's balance is
// checked for sufficient balance.
func (svc *Service) ValidateAccountBalance(ctx context.Context,
	requiredBalance lnwire.MilliSatoshi) error {

	accountID, err := accountFromContext(ctx)
	switch {
	// Something went wrong and needs to bubble up to the caller, possibly
	// rejecting the payment flow.
	case err != nil:
		return err

	// No account was attached to the context, this is just a noop then.
	case accountID == nil:
		return nil
	}
	return svc.checkBalance(*accountID, requiredBalance)
}

// ChargeAccountBalance reads the macaroon from the context and checks
// if it is locked to an account. If it is locked, the account is charged with
// the specified amount, reducing its balance.
func (svc *Service) ChargeAccountBalance(ctx context.Context,
	payment lntypes.Hash, amount lnwire.MilliSatoshi) error {

	accountID, err := accountFromContext(ctx)
	switch {
	// Something went wrong and needs to bubble up to the caller, possibly
	// rejecting the payment flow.
	case err != nil:
		return err

	// No account was attached to the context, this is just a noop then.
	case accountID == nil:
		return nil
	}
	return svc.chargeAccount(*accountID, &payment, amount)
}

func (svc *Service) AddInvoice(ctx context.Context, hash *lntypes.Hash) error {
	accountID, err := accountFromContext(ctx)
	switch {
	// Something went wrong and needs to bubble up to the caller, possibly
	// rejecting the invoice flow.
	case err != nil:
		return err

	// No account was attached to the context, this is just a noop then.
	case accountID == nil:
		return nil
	}
	return svc.addInvoice(*accountID, hash)
}

func accountFromContext(ctx context.Context) (*AccountID, error) {
	// Get the macaroon from the context and see if it is locked to an
	// account.
	mac, err := macaroonFromContext(ctx)
	if err != nil {
		return nil, err
	}
	macaroonAccount := GetCaveatArgOfCondition(mac, CondAccount)
	if len(macaroonAccount) == 0 {
		// There is no condition that locks the macaroon to an account,
		// so there is nothing to check.
		return nil, nil
	}

	// The macaroon is indeed locked to an account. Fetch the account and
	// validate its balance.
	accountIDBytes, err := hex.DecodeString(macaroonAccount)
	if err != nil {
		return nil, err
	}
	var accountID AccountID
	copy(accountID[:], accountIDBytes)
	return &accountID, nil
}

// Close closes the database that underlies the RootKeyStore and AccountStore
// and zeroes the encryption keys.
func (svc *Service) Close() error {
	return svc.AccountStorage.Close()
}
