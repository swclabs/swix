package carts

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

type cache struct {
	cart ICartRepository
}

func useCache(repo ICartRepository) ICartRepository {
	return &cache{
		cart: repo,
	}
}

var _ ICartRepository = (*cache)(nil)

// GetCartByUserID implements ICartRepository.
func (c *cache) GetCartByUserID(ctx context.Context, userId int64, limit int) (*domain.CartSchema, error) {
	return c.cart.GetCartByUserID(ctx, userId, limit)
}

// Insert implements ICartRepository.
func (c *cache) Insert(ctx context.Context, userId int64, inventoryId int64, quantity int64) error {
	return c.cart.Insert(ctx, userId, inventoryId, quantity)
}

// RemoveItem implements ICartRepository.
func (c *cache) RemoveItem(ctx context.Context, inventoryId int64, userId int64) error {
	return c.cart.RemoveItem(ctx, inventoryId, userId)
}
