package addresses

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/pkg/infra/cache"
)

var _ IAddress = (*_cache)(nil)

func useCache(cache cache.ICache, repo IAddress) IAddress {
	return &_cache{
		address: repo,
		cache:   cache,
	}
}

type _cache struct {
	cache   cache.ICache
	address IAddress
}

// GetByUserID implements IAddressRepository.
func (c *_cache) GetByUserID(ctx context.Context, userID int64) ([]entity.Addresses, error) {
	return c.address.GetByUserID(ctx, userID)
}

// Insert implements IAddressRepository.
func (c *_cache) Insert(ctx context.Context, data entity.Addresses) error {
	return c.address.Insert(ctx, data)
}
