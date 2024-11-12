// Package orders implements order repos
package orders

import (
	"context"
	"fmt"

	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/internal/core/domain/model"
	"github.com/swclabs/swipex/pkg/infra/cache"
	"github.com/swclabs/swipex/pkg/lib/crypto"
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

func (c *_Cache) GetLimit(ctx context.Context, limit int) ([]entity.Order, error) {
	return c.orders.GetLimit(ctx, limit)
}

// GetItemByCode implements IOrders.
func (c *_Cache) GetItemByCode(ctx context.Context, orderCode string) ([]model.Order, error) {
	return c.orders.GetItemByCode(ctx, orderCode)
}

// GetByUUID implements IOrders.
func (c *_Cache) GetByUUID(ctx context.Context, orderCode string) (*entity.Order, error) {
	return c.orders.GetByUUID(ctx, orderCode)
}

// Create implements IOrdersRepository.
func (c *_Cache) Create(ctx context.Context, order entity.Order) (int64, error) {
	return c.orders.Create(ctx, order)
}

// GetByUserID implements IOrdersRepository.
func (c *_Cache) GetByUserID(ctx context.Context, userID int64, limit int) ([]entity.Order, error) {
	key := crypto.HashOf(fmt.Sprintf("IOrdersRepository.GetByUserID:%d:%d", userID, limit))
	result, err := cache.GetSlice[entity.Order](ctx, c.cache, key)
	if err != nil {
		result, err = c.orders.GetByUserID(ctx, userID, limit)
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
