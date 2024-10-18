// Package carts carts repos implementation
package carts

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/domain/model"

	"github.com/stretchr/testify/mock"
)

// Mock struct for carts repos
type Mock struct {
	mock.Mock
}

// GetCartInfo implements ICarts.
func (c *Mock) GetCartInfo(ctx context.Context, userID int64) ([]model.Carts, error) {
	panic("unimplemented")
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
