// Package repository
// Author: Duc Hung Ho @kieranhoo
package repository

import (
	"context"
	"log"

	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"
	"swclabs/swipecore/pkg/db/queries"

	"gorm.io/gorm"
)

type Categories struct {
	data *domain.Categories
	conn *gorm.DB
}

func NewCategories() domain.ICategoriesRepository {
	_conn, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	return &Categories{
		data: &domain.Categories{},
		conn: _conn,
	}
}

func (category *Categories) Insert(ctx context.Context, ctg *domain.Categories) error {
	return db.SafeWriteQuery(ctx, category.conn, queries.InsertIntoCategory, ctg.Name, ctg.Description)
}

func (category *Categories) GetLimit(ctx context.Context, limit string) ([]domain.Categories, error) {
	var categories []domain.Categories
	if err := category.conn.WithContext(ctx).Raw(
		queries.SelectCategoryLimit,
		limit,
	).Scan(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
