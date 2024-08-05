package carts

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/pkg/infra/cache"
)

type _cache struct {
	cache cache.ICache
	cart  ICartRepository
}

func useCache(cache cache.ICache, repo ICartRepository) ICartRepository {
	return &_cache{
		cache: cache,
		cart:  repo,
	}
}

var _ ICartRepository = (*_cache)(nil)

// GetCartByUserID implements ICartRepository.
func (c *_cache) GetCartByUserID(ctx context.Context, userID int64, limit int) ([]entity.Carts, error) {
	return c.cart.GetCartByUserID(ctx, userID, limit)
}

// Insert implements ICartRepository.
func (c *_cache) Insert(ctx context.Context, userID int64, inventoryID int64, quantity int64) error {
	return c.cart.Insert(ctx, userID, inventoryID, quantity)
}

// RemoveItem implements ICartRepository.
func (c *_cache) RemoveItem(ctx context.Context, inventoryID int64, userID int64) error {
	return c.cart.RemoveItem(ctx, inventoryID, userID)
}
