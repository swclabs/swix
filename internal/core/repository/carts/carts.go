// Package carts
// Author: Duc Hung Ho @kyeranyo
// Description: cart repository implementation
package carts

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/jackc/pgx/v5"
)

type Carts struct {
	conn *pgx.Conn
}

func New(connection *pgx.Conn) *Carts {
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
