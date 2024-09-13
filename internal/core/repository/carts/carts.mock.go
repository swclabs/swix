// Package carts carts repository implementation
package carts

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"

	"github.com/stretchr/testify/mock"
)

// Mock struct for carts repository
type Mock struct {
	mock.Mock
}

var _ ICarts = (*Mock)(nil)

// NewCartsMock returns a new Mock object
func NewCartsMock() *Mock {
	return &Mock{}
}

// GetCartByUserID implements domain.ICartRepository.
func (c *Mock) GetCartByUserID(_ context.Context, _ int64, _ int) ([]entity.Carts, error) {
	panic("unimplemented")
}

// Insert implements domain.ICartRepository.
func (c *Mock) Insert(_ context.Context, _ entity.Carts) error {
	panic("unimplemented")
}

// RemoveItem implements domain.ICartRepository.
func (c *Mock) RemoveItem(_ context.Context, _ int64) error {
	panic("unimplemented")
}
