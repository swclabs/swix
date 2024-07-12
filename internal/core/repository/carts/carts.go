// Package carts cart repository implementation
package carts

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/infra/db"
)

type Carts struct {
	db db.IDatabase
}

var _ ICartRepository = (*Carts)(nil)

func New(connection db.IDatabase) ICartRepository {
	return useCache(&Carts{
		db: connection,
	})
}

// GetCartByUserID implements domain.ICartRepository.
func (c *Carts) GetCartByUserID(ctx context.Context, userID int64, limit int) (*domain.CartSlices, error) {
	rows, err := c.db.Query(ctx, selectByUserID, userID, limit)
	if err != nil {
		return nil, err
	}
	var cartSchema domain.CartSlices
	cartSchema.UserID = userID

	cartItems, err := db.CollectRows[domain.Carts](rows)
	if err != nil {
		return nil, err
	}
	for _, item := range cartItems {

		cartSchema.Products = append(cartSchema.Products, domain.CartSchema{
			Quantity: item.Quantity,
		})
	}

	return &cartSchema, nil
}

// Insert implements domain.ICartRepository.
func (c *Carts) Insert(ctx context.Context, userID int64, inventoryID int64, quantity int64) error {
	return c.db.SafeWrite(ctx, insertItemToCart,
		userID, inventoryID, quantity,
	)
}

// RemoveItem implements domain.ICartRepository.
func (c *Carts) RemoveItem(ctx context.Context, inventoryID int64, userID int64) error {
	return c.db.SafeWrite(ctx, deleteItem,
		userID, inventoryID,
	)
}
