package carts

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

// ICartRepository implement all method of Carts To access database
type ICartRepository interface {
	// Insert Products to database by productID
	Insert(ctx context.Context, productID int64) error
	// InsertMany insert many Products to database by list of productID
	InsertMany(ctx context.Context, products []int64) error
	// GetCartByUserID is a method get CartInfo from database by userId
	GetCartByUserID(ctx context.Context, userId int64) (*domain.CartInfo, error)
	// RemoveProduct delete Products in Cart
	RemoveProduct(ctx context.Context, productID int64, userId int64) error
}
