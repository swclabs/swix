// Package categories categories repository implementation
package categories

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/pkg/infra/db"
)

// Categories struct for category repository
type Categories struct {
	db db.IDatabase
}

// New creates a new Categories object
func New(conn db.IDatabase) ICategoriesRepository {
	return useCache(&Categories{db: conn})
}

// Insert implements ICategoriesRepository.
func (category *Categories) Insert(ctx context.Context, ctg entity.Categories) error {
	return category.db.SafeWrite(
		ctx, insertIntoCategory, ctg.Name, ctg.Description)
}

// GetLimit implements ICategoriesRepository.
func (category *Categories) GetLimit(ctx context.Context, limit string) ([]entity.Categories, error) {
	rows, err := category.db.Query(ctx, selectCategoryLimit, limit)
	if err != nil {
		return nil, err
	}
	categories, err := db.CollectRows[entity.Categories](rows)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
