// Package products
// Author: Duc Hung Ho @kieranhoo
package products

import (
	"context"
	"encoding/json"
	"strings"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"

	"gorm.io/gorm"
)

type Products struct {
	conn *gorm.DB
}

func New(conn *gorm.DB) *Products {
	return &Products{
		conn: conn,
	}
}

// Insert implements domain.IProductRepository.
func (product *Products) Insert(
	ctx context.Context, prd *domain.Products) (int64, error) {
	return db.SafeWriteQueryReturnId(
		ctx, product.conn, InsertIntoProducts,
		prd.Image, prd.Price, prd.Name, prd.Description,
		prd.SupplierID, prd.CategoryID, prd.Status, prd.Spec,
	)
}

// GetLimit implements domain.IProductRepository.
func (product *Products) GetLimit(
	ctx context.Context, limit int) ([]domain.ProductRes, error) {

	var products []domain.Products
	var productResponse []domain.ProductRes
	if err := product.conn.Table(domain.ProductsTable).
		WithContext(ctx).
		Find(&products).Limit(limit).Error; err != nil {
		return nil, err
	}
	for _, p := range products {
		var spec domain.Specs
		if err := json.Unmarshal([]byte(p.Spec), &spec); err != nil {
			return nil, err // don't find anything, just return empty object
		}
		images := strings.Split(p.Image, ",")
		productResponse = append(productResponse,
			domain.ProductRes{
				ID:          p.ID,
				Price:       p.Price,
				Description: p.Description,
				Name:        p.Name,
				Status:      p.Status,
				Created:     p.Created,
				Image:       images[1:],
				Spec:        spec,
			})
	}
	return productResponse, nil
}

// UploadNewImage implements domain.IProductRepository.
func (product *Products) UploadNewImage(
	ctx context.Context, urlImg string, id int) error {
	return db.SafeWriteQuery(
		ctx, product.conn, UpdateProductImage,
		urlImg, id,
	)
}
