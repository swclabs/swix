package carts

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"
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
func (c *cache) GetCartByUserID(ctx context.Context, userID int64, limit int) ([]entity.Carts, error) {
	return c.cart.GetCartByUserID(ctx, userID, limit)
}

// Insert implements ICartRepository.
func (c *cache) Insert(ctx context.Context, userID int64, inventoryID int64, quantity int64) error {
	return c.cart.Insert(ctx, userID, inventoryID, quantity)
}

// RemoveItem implements ICartRepository.
func (c *cache) RemoveItem(ctx context.Context, inventoryID int64, userID int64) error {
	return c.cart.RemoveItem(ctx, inventoryID, userID)
}
