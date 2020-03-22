package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/btcsuite/btcutil"
	"github.com/guggero/lnd-accounts/acctrpc"
	"github.com/guggero/lnd-accounts/macaroons"
	"github.com/lightningnetwork/lnd/lnwire"
)

type RpcServer struct {
	macService *macaroons.Service
}

// CreateAccount adds an entry to the account database. This entry represents
// an amount of satoshis (account balance) that can be spent using off-chain
// transactions (e.g. paying invoices).
//
// Macaroons can be created to be locked to an account. This makes sure that
// the bearer of the macaroon can only spend at most that amount of satoshis
// through the daemon that has issued the macaroon.
//
// Accounts only assert a maximum amount spendable. Having a certain account
// balance does not guarantee that the node has the channel liquidity to
// actually spend that amount.
func (s *RpcServer) CreateAccount(ctx context.Context,
	req *acctrpc.CreateAccountRequest) (*acctrpc.CreateAccountResponse,
	error) {

	log.Debugf("[createaccount]")

	var (
		balanceMsat    lnwire.MilliSatoshi
		expirationDate time.Time
	)

	// If the expiration date was set, parse it as an unix time stamp.
	// Otherwise we leave it nil to indicate the account has no expiration
	// date.
	if req.ExpirationDate >= 0 {
		expirationDate = time.Unix(req.ExpirationDate, 0)
	}

	// Convert from satoshis to millisatoshis for storage.
	balance := btcutil.Amount(req.AccountBalance)
	balanceMsat = lnwire.NewMSatFromSatoshis(balance)

	// Create the actual account in the macaroon account store.
	account, err := s.macService.NewAccount(
		balanceMsat, expirationDate,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to create account: %v", err)
	}

	// Map the response into the proper response type and return it.
	rpcAccount := &acctrpc.Account{
		Id:             hex.EncodeToString(account.ID[:]),
		InitialBalance: uint64(account.InitialBalance.ToSatoshis()),
		CurrentBalance: uint64(account.CurrentBalance.ToSatoshis()),
		LastUpdate:     account.LastUpdate.Unix(),
		ExpirationDate: int64(0),
	}
	if !account.ExpirationDate.IsZero() {
		rpcAccount.ExpirationDate = account.ExpirationDate.Unix()
	}
	resp := &acctrpc.CreateAccountResponse{
		Account: rpcAccount,
	}
	return resp, nil
}

// ListAccounts returns all accounts that are currently stored in the account
// database.
func (s *RpcServer) ListAccounts(ctx context.Context,
	req *acctrpc.ListAccountsRequest) (*acctrpc.ListAccountsResponse,
	error) {

	log.Debugf("[listaccounts]")

	// Retrieve all accounts from the macaroon account store.
	accounts, err := s.macService.GetAccounts()
	if err != nil {
		return nil, fmt.Errorf("unable to list accounts: %v", err)
	}

	// Map the response into the proper response type and return it.
	rpcAccounts := make([]*acctrpc.Account, len(accounts))
	for i, account := range accounts {
		rpcAccounts[i] = &acctrpc.Account{
			Id: hex.EncodeToString(account.ID[:]),
			InitialBalance: uint64(
				account.InitialBalance.ToSatoshis(),
			),
			CurrentBalance: uint64(
				account.CurrentBalance.ToSatoshis(),
			),
			LastUpdate:     account.LastUpdate.Unix(),
			ExpirationDate: int64(0),
		}
		if !account.ExpirationDate.IsZero() {
			rpcAccounts[i].ExpirationDate =
				account.ExpirationDate.Unix()
		}
	}
	resp := &acctrpc.ListAccountsResponse{
		Accounts: rpcAccounts,
	}
	return resp, nil
}

// RemoveAccount removes the given account from the account database.
func (s *RpcServer) RemoveAccount(ctx context.Context,
	req *acctrpc.RemoveAccountRequest) (*acctrpc.RemoveAccountResponse,
	error) {

	// Account ID is always a hex string, convert it to byte array.
	var accountID = macaroons.AccountID{}
	decoded, err := hex.DecodeString(req.Id)
	if err != nil {
		return nil, err
	}
	copy(accountID[:], decoded)

	// Now remove the account.
	err = s.macService.RemoveAccount(accountID)
	if err != nil {
		return nil, err
	}
	return &acctrpc.RemoveAccountResponse{}, nil
}
