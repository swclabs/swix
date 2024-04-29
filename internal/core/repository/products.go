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

// Insert implements domain.IProductRepository.
func (product *Products) Insert(ctx context.Context, prd *domain.Products) error {
	return db.SafeWriteQuery(
		ctx,
		product.conn,
		queries.InsertIntoProducts,
		prd.Image, prd.Price, prd.Name, prd.Description, prd.Available, prd.SupplierID, prd.CategoryID,
	)
}

// GetLitmit implements domain.IProductRepository.
func (product *Products) GetLitmit(ctx context.Context, limit int) ([]domain.Products, error) {
	var products []domain.Products
	if err := product.conn.Table(domain.ProductsTable).Find(&products).Limit(limit).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// UploadNewImage implements domain.IProductRepository.
func (product *Products) UploadNewImage(ctx context.Context, urlImg string, id int) error {
	return db.SafeWriteQuery(
		ctx,
		product.conn,
		queries.UpdateProductImage,
		urlImg, id,
	)
}
