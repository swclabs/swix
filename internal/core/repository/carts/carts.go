// Package carts
// Author: Duc Hung Ho @kieranhoo
// Description: cart repository implementation
package carts

import (
	"context"
	"log"
	"swclabs/swipecore/internal/core/domain"

	"swclabs/swipecore/pkg/db"

	"gorm.io/gorm"
)

type Carts struct {
	conn *gorm.DB
}

func New() ICartRepository {
	_conn, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	return &Carts{
		conn: _conn,
	}
}

// GetCartByUserID implements domain.ICartRepository.
func (c *Carts) GetCartByUserID(ctx context.Context, userId int64) (*domain.CartInfo, error) {
	panic("unimplemented")
}

// Insert implements domain.ICartRepository.
func (c *Carts) Insert(ctx context.Context, productID int64) error {
	panic("unimplemented")
}

// InsertMany implements domain.ICartRepository.
func (c *Carts) InsertMany(ctx context.Context, products []int64) error {
	panic("unimplemented")
}

// RemoveProduct implements domain.ICartRepository.
func (c *Carts) RemoveProduct(ctx context.Context, productID int64, userId int64) error {
	panic("unimplemented")
}
