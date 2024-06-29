package carts

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

// ICartRepository implement all method of Carts To access database
type ICartRepository interface {
	// Insert Products to database by productID
	Insert(ctx context.Context, userId int64, inventoryId int64, quantity int64) error
	// GetCartByUserID is a method get CartSchema from database by userId
	GetCartByUserID(ctx context.Context, userId int64, limit int) (*domain.CartSchema, error)
	// RemoveItem delete Products in Cart
	RemoveItem(ctx context.Context, inventoryId int64, userId int64) error
}
