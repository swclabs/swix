// Package repo
// Author: Duc Hung Ho @kieranhoo
package repo

import (
	"log"

	"github.com/swclabs/swipe-api/internal/core/domain"
	"github.com/swclabs/swipe-api/pkg/db"
	"github.com/swclabs/swipe-api/pkg/db/queries"
	"gorm.io/gorm"
)

type Products struct {
	data *domain.Products
	conn *gorm.DB
}

func NewProducts() domain.IProductRepository {
	_conn, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	return &Products{
		data: &domain.Products{},
		conn: _conn,
	}
}

func (product *Products) New(prd *domain.Products) error {
	return db.SafeWriteQuery(product.conn,
		queries.InsertIntoProducts,
		prd.Image, prd.Name, prd.Description, prd.Available, prd.SupplierID, prd.CategoryID,
	)
}
