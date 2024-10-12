// Package products implements products repos
package products

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/domain/enum"
	"swclabs/swix/internal/core/domain/model"

	"github.com/stretchr/testify/mock"
)

var _ IProducts = (*Mock)(nil)

// NewProductsMock creates a new Mock.
func NewProductsMock() *Mock {
	return &Mock{}
}

// Mock is a mock for IProductRepository.
type Mock struct {
	mock.Mock
}

// UploadShopImage implements IProducts.
func (p *Mock) UploadShopImage(ctx context.Context, urlImg string, ID int) error {
	panic("unimplemented")
}

// GetByCategory implements IProductRepository.
func (p *Mock) GetByCategory(ctx context.Context, types enum.Category, offset int) ([]model.ProductXCategory, error) {
	args := p.Called(ctx, types, offset)
	return args.Get(0).([]model.ProductXCategory), args.Error(1)
}

// Search implements IProductRepository.
func (p *Mock) Search(_ context.Context, _ string) ([]entity.Products, error) {
	panic("unimplemented")
}

// Update implements IProductRepository.
func (p *Mock) Update(ctx context.Context, product entity.Products) error {
	args := p.Called(ctx, product)
	return args.Error(0)
}

// DeleteByID implements IProductRepository.
func (p *Mock) DeleteByID(ctx context.Context, ID int64) error {
	args := p.Called(ctx, ID)
	return args.Error(0)
}

// GetByID implements IProductRepository.
func (p *Mock) GetByID(ctx context.Context, productID int64) (*entity.Products, error) {
	args := p.Called(ctx, productID)
	return args.Get(0).(*entity.Products), args.Error(1)
}

// GetLimit implements IProductRepository.
func (p *Mock) GetLimit(ctx context.Context, limit int, offset int) ([]entity.Products, error) {
	args := p.Called(ctx, limit, offset)
	return args.Get(0).([]entity.Products), args.Error(1)
}

// Insert implements IProductRepository.
func (p *Mock) Insert(ctx context.Context, prd entity.Products) (int64, error) {
	args := p.Called(ctx, prd)
	return args.Get(0).(int64), args.Error(1)
}

// UploadNewImage implements IProductRepository.
func (p *Mock) UploadNewImage(ctx context.Context, urlImg string, id int) error {
	args := p.Called(ctx, urlImg, id)
	return args.Error(0)
}
