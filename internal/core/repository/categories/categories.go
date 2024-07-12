// Package categories categories repository implementation
package categories

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/infra/db"
)

type Categories struct {
	db db.IDatabase
}

func New(conn db.IDatabase) ICategoriesRepository {
	return useCache(&Categories{db: conn})
}

// Insert implements domain.ICategoriesRepository.
func (category *Categories) Insert(ctx context.Context, ctg domain.Categories) error {
	return category.db.SafeWrite(
		ctx, insertIntoCategory, ctg.Name, ctg.Description)
}

// GetLimit implements domain.ICategoriesRepository.
func (category *Categories) GetLimit(ctx context.Context, limit string) ([]domain.Categories, error) {
	rows, err := category.db.Query(ctx, selectCategoryLimit, limit)
	if err != nil {
		return nil, err
	}
	categories, err := db.CollectRows[domain.Categories](rows)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
