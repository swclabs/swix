package accounts

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/pkg/infra/cache"
)

type _cache struct {
	cache   cache.ICache
	account IAccountRepository
}

var _ IAccountRepository = (*_cache)(nil)

func useCache(cache cache.ICache, repo IAccountRepository) IAccountRepository {
	return &_cache{
		cache:   cache,
		account: repo,
	}
}

// GetByEmail implements IAccountRepository.
func (c *_cache) GetByEmail(ctx context.Context, email string) (*entity.Account, error) {
	return c.account.GetByEmail(ctx, email)
}

// Insert implements IAccountRepository.
func (c *_cache) Insert(ctx context.Context, acc entity.Account) error {
	return c.account.Insert(ctx, acc)
}

// SaveInfo implements IAccountRepository.
func (c *_cache) SaveInfo(ctx context.Context, acc entity.Account) error {
	return c.account.SaveInfo(ctx, acc)
}
