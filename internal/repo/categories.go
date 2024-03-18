// Package repo
// Author: Duc Hung Ho @kieranhoo
package repo

import (
	"log"

	"github.com/swclabs/swipe-server/internal/domain"
	"github.com/swclabs/swipe-server/pkg/db"
	"github.com/swclabs/swipe-server/pkg/db/queries"
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

func (category *Categories) New(ctg *domain.Categories) error {
	return db.SafeWriteQuery(category.conn, queries.InsertIntoCategory, ctg.Name, ctg.Description)
}
