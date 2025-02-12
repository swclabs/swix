// Package accounts account repos implementation
package accounts

import (
	"context"
	"errors"
	"time"

	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/pkg/infra/cache"

	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/pkg/infra/db"
)

var _ = app.Repos(Init)

// Init initializes the Accounts object with database and redis connection
func Init(conn db.IDatabase, cache cache.ICache) IAccounts {
	return useCache(cache, &Accounts{db: conn})
}

// New creates a new Accounts object
func New(conn db.IDatabase) IAccounts {
	return &Accounts{conn}
}

// Accounts struct for account repos
type Accounts struct {
	db db.IDatabase
}

// GetByEmail implements IAccountRepository.
func (account *Accounts) GetByEmail(
	ctx context.Context, email string) (*entity.Account, error) {
	rows, err := account.db.Query(ctx, selectByEmail, email)
	if err != nil {
		return nil, err
	}
	acc, err := db.CollectRow[entity.Account](rows)
	if err != nil {
		return nil, err
	}
	return &acc, nil
}

// Insert implements IAccountRepository.
func (account *Accounts) Insert(ctx context.Context, acc entity.Account) (int64, error) {
	createdAt := time.Now().UTC().Format(time.RFC3339)
	return account.db.SafeWriteReturn(ctx,
		insertIntoAccounts,
		acc.Username, acc.Role, acc.Email, acc.Password,
		createdAt, acc.Type,
	)
}

// SaveInfo implements IAccountRepository.
func (account *Accounts) SaveInfo(
	ctx context.Context, acc entity.Account) error {
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
