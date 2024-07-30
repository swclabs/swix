// Package products implements product
package products

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/internal/core/domain/enum"
	"swclabs/swipecore/internal/core/domain/model"
	"swclabs/swipecore/pkg/infra/cache"
	"swclabs/swipecore/pkg/infra/db"
	"swclabs/swipecore/pkg/lib/errors"
)

var _ IProductRepository = (*Products)(nil)

// New creates a new Products object
func New(conn db.IDatabase) IProductRepository {
	return &Products{
		db: conn,
	}
}

// Init initializes the Products object with database and redis connection
func Init(conn db.IDatabase, cache cache.ICache) IProductRepository {
	return useCache(cache, New(conn))
}

// Products struct for product repository
type Products struct {
	db db.IDatabase
}

// GetByCategory implements IProductRepository.
func (product *Products) GetByCategory(ctx context.Context, types enum.Category, offset int) ([]model.ProductXCategory, error) {
	rows, err := product.db.Query(ctx, selectByCategory, types.String(), offset)
	if err != nil {
		return nil, err
	}
	products, err := db.CollectRows[model.ProductXCategory](rows)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// Search implements IProductRepository.
func (product *Products) Search(ctx context.Context, keyword string) ([]entity.Products, error) {
	rows, err := product.db.Query(ctx, searchByKeyword, keyword)
	if err != nil {
		return nil, errors.Repository("search", err)
	}
	products, err := db.CollectRows[entity.Products](rows)
	if err != nil {
		return nil, errors.Repository("search", err)
	}
	return products, nil
}

// Update implements IProductRepository.
func (product *Products) Update(ctx context.Context, prod entity.Products) error {
	return errors.Repository("safely write data",
		product.db.SafeWrite(ctx, updateByID,
			prod.Name, prod.Price, prod.Description, prod.SupplierID,
			prod.CategoryID, prod.Spec, prod.Status, prod.ID,
		),
	)
}

// DeleteByID implements IProductRepository.
func (product *Products) DeleteByID(ctx context.Context, ID int64) error {
	return errors.Repository(
		"safely write data", product.db.SafeWrite(ctx, deleteByID, ID))
}

// GetByID implements IProductRepository.
func (product *Products) GetByID(ctx context.Context, productID int64) (*entity.Products, error) {
	rows, err := product.db.Query(ctx, selectByID, productID)
	if err != nil {
		return nil, errors.Repository("query", err)
	}
	_product, err := db.CollectOneRow[entity.Products](rows)
	if err != nil {
		return nil, errors.Repository("collect row", err)
	}
	return &_product, nil
}

// Insert implements IProductRepository.
func (product *Products) Insert(ctx context.Context, prd entity.Products) (int64, error) {
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
func (product *Products) GetLimit(ctx context.Context, limit int) ([]entity.Products, error) {
	rows, err := product.db.Query(ctx, selectLimit, limit)
	if err != nil {
		return nil, errors.Repository("query", err)
	}
	products, err := db.CollectRows[entity.Products](rows)
	if err != nil {
		return nil, errors.Repository("collect rows", err)
	}
	return products, nil
}

// UploadNewImage implements IProductRepository.
func (product *Products) UploadNewImage(ctx context.Context, urlImg string, id int) error {
	return errors.Repository("write data", product.db.SafeWrite(
		ctx, updateProductImage,
		urlImg, id,
	))
}
