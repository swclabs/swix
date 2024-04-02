// Package repo
// Author: Duc Hung Ho @kieranhoo
package repo

import (
	"context"
	"log"

	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/pkg/db"
	"swclabs/swipe-api/pkg/db/queries"

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

func (category *Categories) New(ctx context.Context, ctg *domain.Categories) error {
	return db.SafeWriteQuery(ctx, category.conn, queries.InsertIntoCategory, ctg.Name, ctg.Description)
}

func (category *Categories) GetAll(ctx context.Context) ([]domain.Categories, error) {
	panic("not implemented")
}
