package carts

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

var _ ICartRepository = (*Mock)(nil)

func NewCartsMock() *Mock {
	return &Mock{}
}

// GetCartByUserID implements domain.ICartRepository.
func (c *Mock) GetCartByUserID(ctx context.Context, userId int64) (*domain.CartSchema, error) {
	panic("unimplemented")
}

// Insert implements domain.ICartRepository.
func (c *Mock) Insert(ctx context.Context, productID int64) error {
	panic("unimplemented")
}

// InsertMany implements domain.ICartRepository.
func (c *Mock) InsertMany(ctx context.Context, products []int64) error {
	panic("unimplemented")
}

// RemoveItem implements domain.ICartRepository.
func (c *Mock) RemoveItem(ctx context.Context, productID int64, userId int64) error {
	panic("unimplemented")
}
