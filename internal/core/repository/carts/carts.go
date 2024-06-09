// Package carts
// Author: Duc Hung Ho @kieranhoo
// Description: cart repository implementation
package carts

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"gorm.io/gorm"
)

type Carts struct {
	conn *gorm.DB
}

func New(connection *gorm.DB) *Carts {
	return &Carts{
		conn: connection,
	}
}

// GetCartByUserID implements domain.ICartRepository.
func (c *Carts) GetCartByUserID(ctx context.Context, userId int64) (*domain.CartSchema, error) {
	panic("unimplemented")
}

// Insert implements domain.ICartRepository.
func (c *Carts) Insert(ctx context.Context, warehouseId int64) error {
	panic("unimplemented")
}

// InsertMany implements domain.ICartRepository.
func (c *Carts) InsertMany(ctx context.Context, warehouseIds []int64) error {
	panic("unimplemented")
}

// RemoveItem implements domain.ICartRepository.
func (c *Carts) RemoveItem(ctx context.Context, warehouseId int64, userId int64) error {
	panic("unimplemented")
}
