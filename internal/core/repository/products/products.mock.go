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

// DeleteById implements IProductRepository.
func (p *Mock) DeleteById(ctx context.Context, Id int64) error {
	//TODO implement me
	panic("implement me")
}

// GetById implements IProductRepository.
func (p *Mock) GetById(ctx context.Context, productId int64) (*domain.Products, error) {
	panic("unimplemented")
}

// GetLimit implements domain.IProductRepository.
func (p *Mock) GetLimit(ctx context.Context, limit int) ([]domain.ProductRes, error) {
	args := p.Called(ctx, limit)
	return args.Get(0).([]domain.ProductRes), args.Error(1)
}

// Insert implements domain.IProductRepository.
func (p *Mock) Insert(ctx context.Context, prd domain.Products) (int64, error) {
	args := p.Called(ctx, prd)
	return args.Get(0).(int64), args.Error(1)
}

// UploadNewImage implements domain.IProductRepository.
func (p *Mock) UploadNewImage(ctx context.Context, urlImg string, id int) error {
	args := p.Called(ctx, urlImg, id)
	return args.Error(0)
}
