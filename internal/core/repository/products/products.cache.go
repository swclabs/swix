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

// Search implements IProductRepository.
func (c *cache) Search(ctx context.Context, keyword string) ([]domain.Products, error) {
	return c.products.Search(ctx, keyword)
}

// Update implements IProductRepository.
func (c *cache) Update(ctx context.Context, product domain.Products) error {
	return c.products.Update(ctx, product)
}

// DeleteById implements IProductRepository.
func (c *cache) DeleteByID(ctx context.Context, ID int64) error {
	return c.products.DeleteByID(ctx, ID)
}

// GetById implements IProductRepository.
func (c *cache) GetByID(ctx context.Context, productID int64) (*domain.Products, error) {
	return c.products.GetByID(ctx, productID)
}

// GetLimit implements IProductRepository.
func (c *cache) GetLimit(ctx context.Context, limit int) ([]domain.Products, error) {
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
