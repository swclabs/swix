package carts

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/domain/model"
	"swclabs/swix/pkg/infra/cache"
)

func useCache(cache cache.ICache, repo ICarts) ICarts {
	return &_cache{
		cache: cache,
		cart:  repo,
	}
}

var _ ICarts = (*_cache)(nil)

type _cache struct {
	cache cache.ICache
	cart  ICarts
}

// GetCartInfo implements ICarts.
func (c *_cache) GetCartInfo(ctx context.Context, userID int64) ([]model.Carts, error) {
	return c.cart.GetCartInfo(ctx, userID)
}

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
