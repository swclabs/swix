// Package products implements products repository
package products

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

var _ IProductRepository = (*Mock)(nil)

func NewProductsMock() *Mock {
	return &Mock{}
}

// Search implements IProductRepository.
func (p *Mock) Search(_ context.Context, _ string) ([]domain.Products, error) {
	panic("unimplemented")
}

// Update implements IProductRepository.
func (p *Mock) Update(ctx context.Context, product domain.Products) error {
	args := p.Called(ctx, product)
	return args.Error(0)
}

// DeleteById implements IProductRepository.
func (p *Mock) DeleteByID(ctx context.Context, ID int64) error {
	args := p.Called(ctx, ID)
	return args.Error(0)
}

// GetById implements IProductRepository.
func (p *Mock) GetByID(ctx context.Context, productID int64) (*domain.Products, error) {
	args := p.Called(ctx, productID)
	return args.Get(0).(*domain.Products), args.Error(1)
}

// GetLimit implements IProductRepository.
func (p *Mock) GetLimit(ctx context.Context, limit int) ([]domain.Products, error) {
	args := p.Called(ctx, limit)
	return args.Get(0).([]domain.Products), args.Error(1)
}

// Insert implements IProductRepository.
func (p *Mock) Insert(ctx context.Context, prd domain.Products) (int64, error) {
	args := p.Called(ctx, prd)
	return args.Get(0).(int64), args.Error(1)
}

// UploadNewImage implements IProductRepository.
func (p *Mock) UploadNewImage(ctx context.Context, urlImg string, id int) error {
	args := p.Called(ctx, urlImg, id)
	return args.Error(0)
}
