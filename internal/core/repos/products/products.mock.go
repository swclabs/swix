// Package products implements products repos
package products

import (
	"context"

	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/internal/core/domain/enum"
	"github.com/swclabs/swipex/internal/core/domain/model"

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

// Rating implements IProducts.
func (p *Mock) Rating(ctx context.Context, productID int64, rating float64) error {
	panic("unimplemented")
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
func (p *Mock) Search(ctx context.Context, keyword string) ([]entity.Product, error) {
	args := p.Called(ctx, keyword)
	return args.Get(0).([]entity.Product), args.Error(1)
}

// Update implements IProductRepository.
func (p *Mock) Update(ctx context.Context, product entity.Product) error {
	args := p.Called(ctx, product)
	return args.Error(0)
}

// DeleteByID implements IProductRepository.
func (p *Mock) DeleteByID(ctx context.Context, ID int64) error {
	args := p.Called(ctx, ID)
	return args.Error(0)
}

// GetByID implements IProductRepository.
func (p *Mock) GetByID(ctx context.Context, productID int64) (*entity.Product, error) {
	args := p.Called(ctx, productID)
	return args.Get(0).(*entity.Product), args.Error(1)
}

// GetLimit implements IProductRepository.
func (p *Mock) GetLimit(ctx context.Context, limit int, offset int) ([]entity.Product, error) {
	args := p.Called(ctx, limit, offset)
	return args.Get(0).([]entity.Product), args.Error(1)
}

// Insert implements IProductRepository.
func (p *Mock) Insert(ctx context.Context, prd entity.Product) (int64, error) {
	args := p.Called(ctx, prd)
	return args.Get(0).(int64), args.Error(1)
}

// UploadNewImage implements IProductRepository.
func (p *Mock) UploadNewImage(ctx context.Context, urlImg string, id int) error {
	args := p.Called(ctx, urlImg, id)
	return args.Error(0)
}
