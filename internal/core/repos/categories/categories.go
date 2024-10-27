// Package categories categories repos implementation
package categories

import (
	"context"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/pkg/infra/cache"
	"swclabs/swipex/pkg/infra/db"
	"swclabs/swipex/pkg/lib/errors"
)

// New creates a new Categories object
func New(conn db.IDatabase) ICategories {
	return &Categories{db: conn}
}

var _ = app.Repos(Init)

// Init initializes the Categories object with database and redis connection
func Init(conn db.IDatabase, cache cache.ICache) ICategories {
	return useCache(cache, New(conn))
}

// Categories struct for category repos
type Categories struct {
	db db.IDatabase
}

// GetByID implements ICategoriesRepository.
func (category *Categories) GetByID(ctx context.Context, ID int64) (*entity.Category, error) {
	raw, err := category.db.Query(ctx, selectCategoryByID, ID)
	if err != nil {
		return nil, err
	}
	result, err := db.CollectRow[entity.Category](raw)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Insert implements ICategoriesRepository.
func (category *Categories) Insert(ctx context.Context, ctg entity.Category) error {
	return category.db.SafeWrite(
		ctx, insertIntoCategory, ctg.Name, ctg.Description)
}

// GetLimit implements ICategoriesRepository.
func (category *Categories) GetLimit(ctx context.Context, limit string) ([]entity.Category, error) {
	rows, err := category.db.Query(ctx, selectCategoryLimit, limit)
	if err != nil {
		return nil, err
	}
	categories, err := db.CollectRows[entity.Category](rows)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// DeleteByID implements ICategoriesRepository.
func (category *Categories) DeleteByID(ctx context.Context, ID int64) error {
	return errors.Repository(
		"safely write data", category.db.SafeWrite(ctx, deleteByID, ID))
}

// Update implements IProductRepository.
func (category *Categories) Update(ctx context.Context, ctg entity.Category) error {
	return errors.Repository("safely write data",
		category.db.SafeWrite(ctx, updateCategories,
			ctg.ID,
			ctg.Name,
			ctg.Description,
		),
	)
}
