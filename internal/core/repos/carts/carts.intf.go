package carts

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
)

// ICarts implement all method of Carts To access database
type ICarts interface {
	// Insert Products to database by productID
	Insert(ctx context.Context, cart entity.Carts) error
	// GetCartByUserID is a method get CartSlices from database by userId
	GetCartByUserID(ctx context.Context, userID int64, limit int) ([]entity.Carts, error)
	// RemoveItem delete Products in Cart
	RemoveItem(ctx context.Context, ID int64) error
}
