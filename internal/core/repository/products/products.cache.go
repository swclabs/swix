package products

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

type cache struct {
	products IProductRepository
}

var _ IProductRepository = (*cache)(nil)

func useCache(product IProductRepository) IProductRepository {
	return &cache{products: product}
}

// DeleteById implements IProductRepository.
func (c *cache) DeleteById(ctx context.Context, Id int64) error {
	return c.products.DeleteById(ctx, Id)
}

// GetById implements IProductRepository.
func (c *cache) GetById(ctx context.Context, productId int64) (*domain.Products, error) {
	return c.products.GetById(ctx, productId)
}

// GetLimit implements IProductRepository.
func (c *cache) GetLimit(ctx context.Context, limit int) ([]domain.ProductRes, error) {
	return c.products.GetLimit(ctx, limit)
}

// Insert implements IProductRepository.
func (c *cache) Insert(ctx context.Context, prd domain.Products) (int64, error) {
	return c.products.Insert(ctx, prd)
}

// UploadNewImage implements IProductRepository.
func (c *cache) UploadNewImage(ctx context.Context, urlImg string, id int) error {
	return c.products.UploadNewImage(ctx, urlImg, id)
}
