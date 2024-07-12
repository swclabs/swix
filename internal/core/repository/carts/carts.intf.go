package carts

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

// ICartRepository implement all method of Carts To access database
type ICartRepository interface {
	// Insert Products to database by productID
	Insert(ctx context.Context, userID int64, inventoryID int64, quantity int64) error
	// GetCartByUserID is a method get CartSlices from database by userId
	GetCartByUserID(ctx context.Context, userID int64, limit int) (*domain.CartSlices, error)
	// RemoveItem delete Products in Cart
	RemoveItem(ctx context.Context, inventoryID int64, userID int64) error
}
