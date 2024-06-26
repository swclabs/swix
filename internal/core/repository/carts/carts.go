// Package carts
// Author: Duc Hung Ho @kyeranyo
// Description: cart repository implementation
package carts

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"
)

type Carts struct {
	db db.IDatabase
}

var _ ICartRepository = (*Carts)(nil)

func New(connection db.IDatabase) ICartRepository {
	return &Carts{
		db: connection,
	}
}

// GetCartByUserID implements domain.ICartRepository.
func (c *Carts) GetCartByUserID(ctx context.Context, userId int64, limit int) (*domain.CartSchema, error) {
	rows, err := c.db.Query(ctx, selectByUserId, userId, limit)
	if err != nil {
		return nil, err
	}
	var cartSchema domain.CartSchema
	cartSchema.UserId = userId

	cartItems, err := db.CollectRows[domain.Carts](rows)
	if err != nil {
		return nil, err
	}
	for _, item := range cartItems {

		cartSchema.Products = append(cartSchema.Products, domain.CartBodySchema{
			Quantity: item.Quantity,
		})
	}

	return &cartSchema, nil
}

// Insert implements domain.ICartRepository.
func (c *Carts) Insert(ctx context.Context, userId int64, WarehouseId int64, quantity int64) error {
	return c.db.SafeWrite(ctx, insertItemToCart,
		userId, WarehouseId, quantity,
	)
}

// RemoveItem implements domain.ICartRepository.
func (c *Carts) RemoveItem(ctx context.Context, warehouseId int64, userId int64) error {
	return c.db.SafeWrite(ctx, deleteItem,
		userId, warehouseId,
	)
}
