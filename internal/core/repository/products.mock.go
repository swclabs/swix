package repository

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type ProductMock struct {
	mock.Mock
}

var _ domain.IProductRepository = (*ProductMock)(nil)


func NewProductsMock() *ProductMock {
	return &ProductMock{}
}

// GetLitmit implements domain.IProductRepository.
func (p *ProductMock) GetLitmit(ctx context.Context, limit int) ([]domain.Products, error) {
	panic("unimplemented")
}

// Insert implements domain.IProductRepository.
func (p *ProductMock) Insert(ctx context.Context, prd *domain.Products) error {
	panic("unimplemented")
}

// UploadNewImage implements domain.IProductRepository.
func (p *ProductMock) UploadNewImage(ctx context.Context, urlImg string, id int) error {
	panic("unimplemented")
}
