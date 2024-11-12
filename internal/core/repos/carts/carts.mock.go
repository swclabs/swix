// Package carts carts repos implementation
package carts

import (
	"context"

	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/internal/core/domain/model"

	"github.com/stretchr/testify/mock"
)

var _ ICarts = (*Mock)(nil)

// Mock struct for carts repos
type Mock struct {
	mock.Mock
}

// RemoveByItemID implements ICarts.
func (c *Mock) RemoveByItemID(ctx context.Context, userID int64, itemID int64) error {
	panic("unimplemented")
}

// GetCartInfo implements ICarts.
func (c *Mock) GetCartInfo(ctx context.Context, userID int64) ([]model.Carts, error) {
	panic("unimplemented")
}

// NewCartsMock returns a new Mock object
func NewCartsMock() *Mock {
	return &Mock{}
}

// GetCartByUserID implements domain.ICartRepository.
func (c *Mock) GetCartByUserID(_ context.Context, _ int64, _ int) ([]entity.Cart, error) {
	panic("unimplemented")
}

// Insert implements domain.ICartRepository.
func (c *Mock) Insert(_ context.Context, _ entity.Cart) error {
	panic("unimplemented")
}

// RemoveByID implements domain.ICartRepository.
func (c *Mock) RemoveByID(_ context.Context, _ int64) error {
	panic("unimplemented")
}
