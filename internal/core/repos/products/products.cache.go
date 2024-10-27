package products

import (
	"context"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/internal/core/domain/enum"
	"swclabs/swipex/internal/core/domain/model"
	"swclabs/swipex/pkg/infra/cache"
)

var _ IProducts = (*_cache)(nil)

func useCache(cache cache.ICache, product IProducts) IProducts {
	return &_cache{
		products: product,
		cache:    cache,
	}
}

type _cache struct {
	cache    cache.ICache
	products IProducts
}

// Rating implements IProducts.
func (c *_cache) Rating(ctx context.Context, productID int64, rating float64) error {
	return c.products.Rating(ctx, productID, rating)
}

// UploadShopImage implements IProducts.
func (c *_cache) UploadShopImage(ctx context.Context, urlImg string, ID int) error {
	return c.products.UploadShopImage(ctx, urlImg, ID)
}

// GetByCategory implements IProductRepository.
func (c *_cache) GetByCategory(ctx context.Context, types enum.Category, offset int) ([]model.ProductXCategory, error) {
	return c.products.GetByCategory(ctx, types, offset)
}

// Search implements IProductRepository.
func (c *_cache) Search(ctx context.Context, keyword string) ([]entity.Product, error) {
	return c.products.Search(ctx, keyword)
}

// Update implements IProductRepository.
func (c *_cache) Update(ctx context.Context, product entity.Product) error {
	return c.products.Update(ctx, product)
}

// DeleteByID implements IProductRepository.
func (c *_cache) DeleteByID(ctx context.Context, ID int64) error {
	return c.products.DeleteByID(ctx, ID)
}

// GetByID implements IProductRepository.
func (c *_cache) GetByID(ctx context.Context, productID int64) (*entity.Product, error) {
	return c.products.GetByID(ctx, productID)
}

// GetLimit implements IProductRepository.
func (c *_cache) GetLimit(ctx context.Context, limit int, offset int) ([]entity.Product, error) {
	return c.products.GetLimit(ctx, limit, offset)
}

// Insert implements IProductRepository.
func (c *_cache) Insert(ctx context.Context, prd entity.Product) (int64, error) {
	return c.products.Insert(ctx, prd)
}

// UploadNewImage implements IProductRepository.
func (c *_cache) UploadNewImage(ctx context.Context, urlImg string, id int) error {
	return c.products.UploadNewImage(ctx, urlImg, id)
}
