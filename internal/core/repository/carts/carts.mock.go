// Package carts carts repository implementation
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
func (c *Mock) GetCartByUserID(_ context.Context, _ int64, _ int) (*domain.CartSlices, error) {
	panic("unimplemented")
}

// Insert implements domain.ICartRepository.
func (c *Mock) Insert(_ context.Context, _ int64, _ int64, _ int64) error {
	panic("unimplemented")
}

// RemoveItem implements domain.ICartRepository.
func (c *Mock) RemoveItem(_ context.Context, _ int64, _ int64) error {
	panic("unimplemented")
}
