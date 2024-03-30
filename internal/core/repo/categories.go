// Package repo
// Author: Duc Hung Ho @kieranhoo
package repo

import (
	"log"

	"gorm.io/gorm"
	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/pkg/db"
	"swclabs/swipe-api/pkg/db/queries"
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

func (category *Categories) New(ctg *domain.Categories) error {
	return db.SafeWriteQuery(category.conn, queries.InsertIntoCategory, ctg.Name, ctg.Description)
}
