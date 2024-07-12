// Package accounts account repository implementation
package accounts

import (
	"context"
	"errors"
	"time"

	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/infra/db"
)

type Accounts struct {
	db db.IDatabase
}

func New(conn db.IDatabase) IAccountRepository {
	return useCache(&Accounts{conn})
}

// GetByEmail implements IAccountRepository.
func (account *Accounts) GetByEmail(
	ctx context.Context, email string) (*domain.Account, error) {
	rows, err := account.db.Query(ctx, selectByEmail, email)
	if err != nil {
		return nil, err
	}
	acc, err := db.CollectOneRow[domain.Account](rows)
	if err != nil {
		return nil, err
	}
	return &acc, nil
}

// Insert implements IAccountRepository.
func (account *Accounts) Insert(
	ctx context.Context, acc domain.Account) error {
	createdAt := time.Now().UTC().Format(time.RFC3339)
	return account.db.SafeWrite(ctx,
		insertIntoAccounts,
		acc.Username, acc.Role, acc.Email, acc.Password,
		createdAt, acc.Type,
	)
}

// SaveInfo implements IAccountRepository.
func (account *Accounts) SaveInfo(
	ctx context.Context, acc domain.Account) error {
	if acc.Email == "" {
		return errors.New("missing key: email ")
	}
	if acc.Username != "" {
		if err := account.db.SafeWrite(ctx, updateAccountsUsername,
			acc.Username, acc.Email); err != nil {
			return err
		}

	}
	if acc.Password != "" {
		if err := account.db.SafeWrite(ctx, updateAccountsPassword,
			acc.Password, acc.Email); err != nil {
			return err
		}
	}
	if acc.Role != "" {
		if err := account.db.SafeWrite(ctx, updateAccountsRole,
			acc.Role, acc.Email); err != nil {
			return err
		}
	}
	return nil
}
