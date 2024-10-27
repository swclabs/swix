package carts

import (
	"context"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/internal/core/domain/model"
)

// ICarts implement all method of Carts To access database
type ICarts interface {
	// Insert Products to database by productID
	Insert(ctx context.Context, cart entity.Cart) error

	// GetCartByUserID is a method get CartSlices from database by userId
	GetCartByUserID(ctx context.Context, userID int64, limit int) ([]entity.Cart, error)

	GetCartInfo(ctx context.Context, userID int64) ([]model.Carts, error)

	// RemoveByID delete Products in Cart
	RemoveByID(ctx context.Context, ID int64) error

	RemoveByItemID(ctx context.Context, userID, itemID int64) error
}
