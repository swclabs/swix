package addresses

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/pkg/infra/cache"
)

type _cache struct {
	cache   cache.ICache
	address IAddressRepository
}

var _ IAddressRepository = (*_cache)(nil)

func useCache(cache cache.ICache, repo IAddressRepository) IAddressRepository {
	return &_cache{
		address: repo,
		cache:   cache,
	}
}

// Insert implements IAddressRepository.
func (c *_cache) Insert(ctx context.Context, data entity.Addresses) error {
	return c.address.Insert(ctx, data)
}
