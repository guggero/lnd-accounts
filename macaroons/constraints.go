package macaroons

import (
	"context"
	"encoding/hex"
	"fmt"

	"gopkg.in/macaroon-bakery.v2/bakery/checkers"
	"gopkg.in/macaroon.v2"
)

const (
	// CondAccount is the caveat condition that binds a macaroon to a
	// certain account.
	CondAccount = "account"
)

// GetCaveatArgOfCondition parses the caveats of a macaroon and returns the
// argument of the caveat that matches the condition name given.
func GetCaveatArgOfCondition(mac *macaroon.Macaroon, cond string) string {
	for _, caveat := range mac.Caveats() {
		caveatCond, arg, err := checkers.ParseCaveat(string(caveat.Id))
		if err != nil {
			// If we can't parse the caveat, it probably doesn't
			// concern us.
			continue
		}
		if caveatCond == cond {
			return arg
		}
	}
	return ""
}

// AccountLockConstraint locks a macaroon to a specific account. So every action
// that sends funds requires this specific account to have enough balance to
// perform the action. Once the action was successful, the sent amount is
// subtracted from the account's balance.
func AccountLockConstraint(accountID AccountID) func(*macaroon.Macaroon) error {
	return func(mac *macaroon.Macaroon) error {
		caveat := checkers.Condition(
			CondAccount, hex.EncodeToString(accountID[:]),
		)
		return mac.AddFirstPartyCaveat([]byte(caveat))
	}
}

// AccountLockChecker checks that the account the macaroon is locked to exists
// and has not yet expired. The actual balance check happens in the rpcserver.
func AccountLockChecker(service *Service) (string, checkers.Func) {
	return CondAccount, func(ctx context.Context, cond, arg string) error {
		id, err := hex.DecodeString(arg)
		if err != nil {
			return fmt.Errorf("invalid account id: %v", err)
		}
		if len(id) != AccountIDLen {
			return fmt.Errorf("invalid account id length")
		}
		var accountID AccountID
		copy(accountID[:], id)
		account, err := service.GetAccount(accountID)
		if err != nil {
			return err
		}

		if account.HasExpired() {
			return ErrAccExpired
		}
		return nil
	}
}
