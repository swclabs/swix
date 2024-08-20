package categories

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/pkg/infra/cache"
)

type _cache struct {
	cache    cache.ICache
	category ICategoriesRepository
}

var _ ICategoriesRepository = (*_cache)(nil)

func useCache(cache cache.ICache, repo ICategoriesRepository) ICategoriesRepository {
	return &_cache{
		category: repo,
		cache:    cache,
	}
}

// GetByID implements ICategoriesRepository.
func (c *_cache) GetByID(ctx context.Context, ID int64) (*entity.Categories, error) {
	return c.category.GetByID(ctx, ID)
}

// GetLimit implements ICategoriesRepository.
func (c *_cache) GetLimit(ctx context.Context, limit string) ([]entity.Categories, error) {
	return c.category.GetLimit(ctx, limit)
}

// Insert implements ICategoriesRepository.
func (c *_cache) Insert(ctx context.Context, ctg entity.Categories) error {
	return c.category.Insert(ctx, ctg)
}

func (c *_cache) DeleteByID(ctx context.Context, ID int64) error {
	return c.category.DeleteByID(ctx, ID)
}

// Update implements IProductRepository.
func (c *_cache) Update(ctx context.Context, ctg entity.Categories) error {
	return c.category.Update(ctx, ctg)
}
