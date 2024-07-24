// Package carts cart repository implementation
package carts

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/pkg/infra/db"
)

// Carts struct for cart repository
type Carts struct {
	db db.IDatabase
}

var _ ICartRepository = (*Carts)(nil)

// New creates a new Carts object
func New(connection db.IDatabase) ICartRepository {
	return useCache(&Carts{
		db: connection,
	})
}

// GetCartByUserID implements domain.ICartRepository.
func (c *Carts) GetCartByUserID(ctx context.Context, userID int64, limit int) ([]entity.Carts, error) {
	rows, err := c.db.Query(ctx, selectByUserID, userID, limit)
	if err != nil {
		return nil, err
	}

	cartItems, err := db.CollectRows[entity.Carts](rows)
	if err != nil {
		return nil, err
	}

	return cartItems, nil
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
