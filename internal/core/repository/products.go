// Package repository
// Author: Duc Hung Ho @kieranhoo
package repository

import (
	"context"
	"encoding/json"
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
func (product *Products) Insert(ctx context.Context, prd *domain.Products) (int64, error) {
	return db.SafeWriteQueryReturnId(
		ctx,
		product.conn,
		queries.InsertIntoProducts,
		prd.Image, prd.Price, prd.Name, prd.Description,
		prd.SupplierID, prd.CategoryID, prd.Status, prd.Spec,
	)
}

// GetLitmit implements domain.IProductRepository.
func (product *Products) GetLitmit(ctx context.Context, limit int) ([]domain.ProductResponse, error) {
	var products []domain.Products
	var productResponse []domain.ProductResponse
	if err := product.conn.Table(domain.ProductsTable).Find(&products).Limit(limit).Error; err != nil {
		return nil, err
	}
	for _, p := range products {
		var spec domain.Specifications
		if err := json.Unmarshal([]byte(p.Spec), &spec); err != nil {
			return nil, err
		}
		productResponse = append(productResponse, domain.ProductResponse{
			ID:          p.ID,
			Image:       p.Image,
			Price:       p.Price,
			Description: p.Description,
			Name:        p.Name,
			Status:      p.Status,
			Spec:        spec,
		})
	}
	return productResponse, nil
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
