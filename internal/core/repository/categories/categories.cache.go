package categories

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/pkg/infra/cache"
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

// GetLimit implements ICategoriesRepository.
func (c *_cache) GetLimit(ctx context.Context, limit string) ([]entity.Categories, error) {
	return c.category.GetLimit(ctx, limit)
}

// Insert implements ICategoriesRepository.
func (c *_cache) Insert(ctx context.Context, ctg entity.Categories) error {
	return c.category.Insert(ctx, ctg)
}
