// Package orders implements order repository
package orders

import (
	"context"
	"fmt"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/pkg/infra/cache"
	"swclabs/swix/pkg/lib/crypto"
)

func useCache(cache cache.ICache, orders IOrders) IOrders {
	return &_Cache{
		cache:  cache,
		orders: orders,
	}
}

type _Cache struct {
	cache  cache.ICache
	orders IOrders
}

// Create implements IOrdersRepository.
func (c *_Cache) Create(ctx context.Context, order entity.Orders) (int64, error) {
	return c.orders.Create(ctx, order)
}

// Get implements IOrdersRepository.
func (c *_Cache) Get(ctx context.Context, userID int64, limit int) ([]entity.Orders, error) {
	key := crypto.HashOf(fmt.Sprintf("IOrdersRepository.Get:%d:%d", userID, limit))
	result, err := cache.GetSlice[entity.Orders](ctx, c.cache, key)
	if err != nil {
		result, err = c.orders.Get(ctx, userID, limit)
		if err != nil {
			return nil, err
		}
		if err := cache.Set(ctx, c.cache, key, result); err != nil {
			// write log for error
			return result, err
		}
	}
	return result, nil
}

// GetProductByOrderID implements IOrdersRepository.
func (c *_Cache) GetProductByOrderID(ctx context.Context, orderID int64) ([]entity.ProductInOrder, error) {
	key := crypto.HashOf(fmt.Sprintf("IOrdersRepository.GetProductByOrderID:%d", orderID))
	result, err := cache.GetSlice[entity.ProductInOrder](ctx, c.cache, key)
	if err != nil {
		result, err = c.orders.GetProductByOrderID(ctx, orderID)
		if err != nil {
			return nil, err
		}
		if err := cache.Set(ctx, c.cache, key, result); err != nil {
			// write log for error
			return result, err
		}
	}
	return result, nil
}

// InsertProduct implements IOrdersRepository.
func (c *_Cache) InsertProduct(ctx context.Context, product entity.ProductInOrder) error {
	return c.orders.InsertProduct(ctx, product)
}
