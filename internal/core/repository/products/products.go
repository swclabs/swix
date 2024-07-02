// Package products
// Author: Duc Hung Ho @kyeranyo
package products

import (
	"context"
	"encoding/json"
	"strings"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/errors"
	"swclabs/swipecore/pkg/db"
	"swclabs/swipecore/pkg/utils"
)

type Products struct {
	db db.IDatabase
}

var _ IProductRepository = (*Products)(nil)

func New(conn db.IDatabase) IProductRepository {
	return useCache(&Products{
		db: conn,
	})
}

// Search implements IProductRepository.
func (product *Products) Search(ctx context.Context, keyword string) ([]domain.Products, error) {
	rows, err := product.db.Query(ctx, searchByKeyword, keyword)
	if err != nil {
		return nil, errors.Repository("search", err)
	}
	products, err := db.CollectRows[domain.Products](rows)
	if err != nil {
		return nil, errors.Repository("search", err)
	}
	return products, nil
}

// Update implements IProductRepository.
func (product *Products) Update(ctx context.Context, prod domain.Products) error {
	return errors.Repository("safely write data",
		product.db.SafeWrite(ctx, updateById,
			prod.Name, prod.Price, prod.Description, prod.SupplierID,
			prod.CategoryID, prod.Spec, prod.Status, prod.ID,
		),
	)
}

// DeleteById implements IProductRepository.
func (product *Products) DeleteById(ctx context.Context, Id int64) error {
	return errors.Repository(
		"safely write data", product.db.SafeWrite(ctx, deleteById, Id))
}

// GetById implements IProductRepository.
func (product *Products) GetById(ctx context.Context, productId int64) (*domain.Products, error) {
	rows, err := product.db.Query(ctx, selectById, productId)
	if err != nil {
		return nil, errors.Repository("query", err)
	}
	_product, err := db.CollectOneRow[domain.Products](rows)
	if err != nil {
		return nil, errors.Repository("collect row", err)
	}
	return &_product, nil
}

// Insert implements IProductRepository.
func (product *Products) Insert(ctx context.Context, prd domain.Products) (int64, error) {
	id, err := product.db.SafeWriteReturn(
		ctx, insertIntoProducts,
		prd.Image, prd.Price, prd.Name, prd.Description,
		prd.SupplierID, prd.CategoryID, prd.Status, prd.Spec,
	)
	if err != nil {
		return -1, errors.Repository("write data", err)
	}
	return id, nil
}

// GetLimit implements IProductRepository.
func (product *Products) GetLimit(ctx context.Context, limit int) ([]domain.ProductSchema, error) {

	var productResponse []domain.ProductSchema

	rows, err := product.db.Query(ctx, selectLimit, limit)
	if err != nil {
		return nil, errors.Repository("query", err)
	}

	products, err := db.CollectRows[domain.Products](rows)
	if err != nil {
		return nil, errors.Repository("collect rows", err)
	}

	for _, p := range products {
		var spec domain.Specs
		if err := json.Unmarshal([]byte(p.Spec), &spec); err != nil {
			// don't find anything, just return empty object
			return nil, errors.Repository("json", err)
		}
		images := strings.Split(p.Image, ",")
		productResponse = append(productResponse,
			domain.ProductSchema{
				ID:          p.ID,
				Price:       p.Price,
				Description: p.Description,
				Name:        p.Name,
				Status:      p.Status,
				Created:     utils.HanoiTimezone(p.Created),
				Image:       images[1:],
				Spec:        spec,
			})
	}
	return productResponse, nil
}

// UploadNewImage implements IProductRepository.
func (product *Products) UploadNewImage(ctx context.Context, urlImg string, id int) error {
	return errors.Repository("write data", product.db.SafeWrite(
		ctx, updateProductImage,
		urlImg, id,
	))
}
