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
	"time"
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

// Update implements IProductRepository.
func (product *Products) Update(ctx context.Context, prod domain.Products) error {
	return errors.Repository("safely write data",
		product.db.SafeWrite(ctx, updateById, 
			prod.Name, prod.Price,prod.Description, prod.SupplierID, 
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

// Insert implements domain.IProductRepository.
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

// GetLimit implements domain.IProductRepository.
func (product *Products) GetLimit(ctx context.Context, limit int) ([]domain.ProductRes, error) {

	var productResponse []domain.ProductRes

	rows, err := product.db.Query(ctx, selectLimit, limit)
	if err != nil {
		return nil, errors.Repository("query", err)
	}

	products, err := db.CollectRows[domain.Products](rows)
	if err != nil {
		return nil, errors.Repository("collect rows", err)
	}

	for _, p := range products {
		var (
			spec   domain.Specs
			isSpec = true
		)
		if err := json.Unmarshal([]byte(p.Spec), &spec); err != nil {
			// don't find anything, just return empty object
			return nil, errors.Repository("json", err)
		}
		if spec.Screen == "" {
			isSpec = false
		}
		images := strings.Split(p.Image, ",")
		productResponse = append(productResponse,
			domain.ProductRes{
				ID:          p.ID,
				Price:       p.Price,
				Description: p.Description,
				Name:        p.Name,
				Status:      p.Status,
				Created:     p.Created.In(time.FixedZone("GMT+7", 7*60*60)).Format(time.DateTime),
				Image:       images[1:],
				IsSpec:      isSpec,
				Spec:        spec,
			})
	}
	return productResponse, nil
}

// UploadNewImage implements domain.IProductRepository.
func (product *Products) UploadNewImage(ctx context.Context, urlImg string, id int) error {
	return errors.Repository("write data", product.db.SafeWrite(
		ctx, updateProductImage,
		urlImg, id,
	))
}
