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
func (c *Mock) GetCartByUserID(ctx context.Context, userId int64, limit int) (*domain.CartSchema, error) {
	panic("unimplemented")
}

// Insert implements domain.ICartRepository.
func (c *Mock) Insert(ctx context.Context, userId int64, inventoryId int64, quantity int64) error {
	panic("unimplemented")
}

// RemoveItem implements domain.ICartRepository.
func (c *Mock) RemoveItem(ctx context.Context, inventoryId int64, userId int64) error {
	panic("unimplemented")
}
