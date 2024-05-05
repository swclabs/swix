package repository

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type CartsMock struct {
	mock.Mock
}

var _ domain.ICartRepository = (*CartsMock)(nil)

func NewCartsMock() *CartsMock {
	return &CartsMock{}
}

// GetCartByUserID implements domain.ICartRepository.
func (c *CartsMock) GetCartByUserID(ctx context.Context, userId int64) (*domain.CartInfo, error) {
	panic("unimplemented")
}

// Insert implements domain.ICartRepository.
func (c *CartsMock) Insert(ctx context.Context, productID int64) error {
	panic("unimplemented")
}

// InsertMany implements domain.ICartRepository.
func (c *CartsMock) InsertMany(ctx context.Context, products []int64) error {
	panic("unimplemented")
}

// RemoveProduct implements domain.ICartRepository.
func (c *CartsMock) RemoveProduct(ctx context.Context, productID int64, userId int64) error {
	panic("unimplemented")
}
