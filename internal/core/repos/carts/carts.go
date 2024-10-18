// Package carts cart repos implementation
package carts

import (
	"context"
	"swclabs/swix/app"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/domain/model"
	"swclabs/swix/pkg/infra/cache"
	"swclabs/swix/pkg/infra/db"
)

var _ ICarts = (*Carts)(nil)
var _ = app.Repos(Init)

// Init initializes the Carts object with database and redis connection
func Init(connection db.IDatabase, cache cache.ICache) ICarts {
	return useCache(cache, New(connection))
}

// New creates a new Carts object
func New(connection db.IDatabase) ICarts {
	return &Carts{
		db: connection,
	}
}

// Carts struct for cart repos
type Carts struct {
	db db.IDatabase
}

// GetCartInfo implements ICarts.
func (c *Carts) GetCartInfo(ctx context.Context, userID int64) ([]model.Carts, error) {
	rows, err := c.db.Query(ctx, getCartInfo, userID)
	if err != nil {
		return nil, err
	}
	cartItems, err := db.CollectRows[model.Carts](rows)
	if err != nil {
		return nil, err
	}
	return cartItems, nil
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
func (c *Carts) Insert(ctx context.Context, cart entity.Carts) error {
	return c.db.SafeWrite(ctx, insertItemToCart,
		cart.UserID, cart.InventoryID, cart.Quantity,
	)
}

// RemoveItem implements domain.ICartRepository.
func (c *Carts) RemoveItem(ctx context.Context, ID int64) error {
	return c.db.SafeWrite(ctx, deleteItem, ID)
}
