package categories

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

type cache struct {
	category ICategoriesRepository
}

var _ ICategoriesRepository = (*cache)(nil)

func useCache(repo ICategoriesRepository) ICategoriesRepository {
	return &cache{
		category: repo,
	}
}

// GetLimit implements ICategoriesRepository.
func (c *cache) GetLimit(ctx context.Context, limit string) ([]domain.Categories, error) {
	return c.category.GetLimit(ctx, limit)
}

// Insert implements ICategoriesRepository.
func (c *cache) Insert(ctx context.Context, ctg domain.Categories) error {
	return c.category.Insert(ctx, ctg)
}
