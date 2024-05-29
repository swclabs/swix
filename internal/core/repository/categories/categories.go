// Package repository
// Author: Duc Hung Ho @kieranhoo
package categories

import (
	"context"
	"log"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"

	"gorm.io/gorm"
)

type Categories struct {
	data *domain.Categories
	conn *gorm.DB
}

func NewCategories() ICategoriesRepository {
	_conn, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	return &Categories{
		data: &domain.Categories{},
		conn: _conn,
	}
}

// Insert implements domain.ICategoriesRepository.
func (category *Categories) Insert(ctx context.Context, ctg *domain.Categories) error {
	return db.SafeWriteQuery(ctx, category.conn, InsertIntoCategory, ctg.Name, ctg.Description)
}

// GetLimit implements domain.ICategoriesRepository.
func (category *Categories) GetLimit(ctx context.Context, limit string) ([]domain.Categories, error) {
	var categories []domain.Categories
	if err := category.conn.WithContext(ctx).Raw(
		SelectCategoryLimit,
		limit,
	).Scan(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
