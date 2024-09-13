package carts

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/pkg/infra/cache"
)

type _cache struct {
	cache cache.ICache
	cart  ICarts
}

func useCache(cache cache.ICache, repo ICarts) ICarts {
	return &_cache{
		cache: cache,
		cart:  repo,
	}
}

var _ ICarts = (*_cache)(nil)

// GetCartByUserID implements ICartRepository.
func (c *_cache) GetCartByUserID(ctx context.Context, userID int64, limit int) ([]entity.Carts, error) {
	return c.cart.GetCartByUserID(ctx, userID, limit)
}

// Insert implements ICartRepository.
func (c *_cache) Insert(ctx context.Context, cart entity.Carts) error {
	return c.cart.Insert(ctx, cart)
}

// RemoveItem implements ICartRepository.
func (c *_cache) RemoveItem(ctx context.Context, ID int64) error {
	return c.cart.RemoveItem(ctx, ID)
}
