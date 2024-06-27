// Package products
// Author: Duc Hung Ho @kyeranyo
package products

import (
	"context"
	"encoding/json"
	"strings"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"

	"github.com/jackc/pgx/v5"
)

type Products struct {
	db db.IDatabase
}

var _ IProductRepository = (*Products)(nil)

func New(conn db.IDatabase) IProductRepository {
	return &Products{
		db: conn,
	}
}

// DeleteById implements IProductRepository.
func (product *Products) DeleteById(ctx context.Context, Id int64) error {
	return product.db.SafeWrite(ctx, deleteById, Id)
}

// GetById implements IProductRepository.
func (product *Products) GetById(ctx context.Context, productId int64) (*domain.Products, error) {
	rows, err := product.db.Query(ctx, selectById, productId)
	if err != nil {
		return nil, err
	}
	_product, err := pgx.CollectOneRow[domain.Products](rows, pgx.RowToStructByName[domain.Products])
	if err != nil {
		return nil, err
	}
	return &_product, nil
}

// Insert implements domain.IProductRepository.
func (product *Products) Insert(
	ctx context.Context, prd domain.Products) (int64, error) {
	return product.db.SafeWriteReturn(
		ctx, insertIntoProducts,
		prd.Image, prd.Price, prd.Name, prd.Description,
		prd.SupplierID, prd.CategoryID, prd.Status, prd.Spec,
	)
}

// GetLimit implements domain.IProductRepository.
func (product *Products) GetLimit(
	ctx context.Context, limit int) ([]domain.ProductRes, error) {

	var productResponse []domain.ProductRes

	rows, err := product.db.Query(ctx, selectLimit, limit)
	if err != nil {
		return nil, err
	}

	products, err := pgx.CollectRows[domain.Products](rows, pgx.RowToStructByName[domain.Products])
	if err != nil {
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
	return product.db.SafeWrite(
		ctx, updateProductImage,
		urlImg, id,
	)
}
