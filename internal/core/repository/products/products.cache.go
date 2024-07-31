package products

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/internal/core/domain/enum"
	"swclabs/swipecore/internal/core/domain/model"
	"swclabs/swipecore/pkg/infra/cache"
)

var _ IProductRepository = (*_cache)(nil)

func useCache(cache cache.ICache, product IProductRepository) IProductRepository {
	return &_cache{
		products: product,
		cache:    cache,
	}
}

type _cache struct {
	cache    cache.ICache
	products IProductRepository
}

// GetByCategory implements IProductRepository.
func (c *_cache) GetByCategory(ctx context.Context, types enum.Category, offset int) ([]model.ProductXCategory, error) {
	return c.products.GetByCategory(ctx, types, offset)
}

// Search implements IProductRepository.
func (c *_cache) Search(ctx context.Context, keyword string) ([]entity.Products, error) {
	return c.products.Search(ctx, keyword)
}

// Update implements IProductRepository.
func (c *_cache) Update(ctx context.Context, product entity.Products) error {
	return c.products.Update(ctx, product)
}

// DeleteById implements IProductRepository.
func (c *_cache) DeleteByID(ctx context.Context, ID int64) error {
	return c.products.DeleteByID(ctx, ID)
}

// GetById implements IProductRepository.
func (c *_cache) GetByID(ctx context.Context, productID int64) (*entity.Products, error) {
	return c.products.GetByID(ctx, productID)
}

// GetLimit implements IProductRepository.
func (c *_cache) GetLimit(ctx context.Context, limit int) ([]entity.Products, error) {
	return c.products.GetLimit(ctx, limit)
}

// Insert implements IProductRepository.
func (c *_cache) Insert(ctx context.Context, prd entity.Products) (int64, error) {
	return c.products.Insert(ctx, prd)
}

// UploadNewImage implements IProductRepository.
func (c *_cache) UploadNewImage(ctx context.Context, urlImg string, id int) error {
	return c.products.UploadNewImage(ctx, urlImg, id)
}
