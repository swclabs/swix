// Package categories
// Author: Duc Hung Ho @kyeranyo
package categories

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"

	"github.com/jackc/pgx/v5"
)

type Categories struct {
	db db.IDatabase
}

func New(conn db.IDatabase) ICategoriesRepository {
	return &Categories{db: conn}
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
	categories, err := pgx.CollectRows[domain.Categories](rows, pgx.RowToStructByName[domain.Categories])
	if err != nil {
		return nil, err
	}
	return categories, nil
}
