package products

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type ProductMock struct {
	mock.Mock
}

var _ IProductRepository = (*ProductMock)(nil)

func NewProductsMock() *ProductMock {
	return &ProductMock{}
}

// GetLimit implements domain.IProductRepository.
func (p *ProductMock) GetLimit(ctx context.Context, limit int) ([]domain.ProductRes, error) {
	panic("unimplemented")
}

// Insert implements domain.IProductRepository.
func (p *ProductMock) Insert(ctx context.Context, prd *domain.Products) (int64, error) {
	panic("unimplemented")
}

// UploadNewImage implements domain.IProductRepository.
func (p *ProductMock) UploadNewImage(ctx context.Context, urlImg string, id int) error {
	panic("unimplemented")
}
