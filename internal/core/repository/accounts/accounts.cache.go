package accounts

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"
)

type cache struct {
	account IAccountRepository
}

var _ IAccountRepository = (*cache)(nil)

func useCache(repo IAccountRepository) IAccountRepository {
	return &cache{account: repo}
}

// GetByEmail implements IAccountRepository.
func (c *cache) GetByEmail(ctx context.Context, email string) (*entity.Account, error) {
	return c.account.GetByEmail(ctx, email)
}

// Insert implements IAccountRepository.
func (c *cache) Insert(ctx context.Context, acc entity.Account) error {
	return c.account.Insert(ctx, acc)
}

// SaveInfo implements IAccountRepository.
func (c *cache) SaveInfo(ctx context.Context, acc entity.Account) error {
	return c.account.SaveInfo(ctx, acc)
}
