package carts

import (
	"context"

	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/internal/core/domain/model"
	"github.com/swclabs/swipex/pkg/infra/cache"
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

// RemoveByItemID implements ICarts.
func (c *_cache) RemoveByItemID(ctx context.Context, userID int64, itemID int64) error {
	return c.cart.RemoveByItemID(ctx, userID, itemID)
}

// GetCartInfo implements ICarts.
func (c *_cache) GetCartInfo(ctx context.Context, userID int64) ([]model.Carts, error) {
	return c.cart.GetCartInfo(ctx, userID)
}

// GetCartByUserID implements ICartRepository.
func (c *_cache) GetCartByUserID(ctx context.Context, userID int64, limit int) ([]entity.Cart, error) {
	return c.cart.GetCartByUserID(ctx, userID, limit)
}

// Insert implements ICartRepository.
func (c *_cache) Insert(ctx context.Context, cart entity.Cart) error {
	return c.cart.Insert(ctx, cart)
}

// RemoveByID implements ICartRepository.
func (c *_cache) RemoveByID(ctx context.Context, ID int64) error {
	return c.cart.RemoveByID(ctx, ID)
}
