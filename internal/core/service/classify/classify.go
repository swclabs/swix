// Package classify implements classify service include category and supplier
package classify

import (
	"context"
	"log"
	"swclabs/swix/app"

	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/repos/categories"
	"swclabs/swix/internal/core/repos/suppliers"
	"swclabs/swix/pkg/infra/db"
)

var _ = app.Service(New)

// New creates a new Classify object
func New(
	category categories.ICategories,
	supplier suppliers.ISuppliers,
) IClassify {
	return &Classify{
		Category: category,
		Supplier: supplier,
	}
}

// Classify struct for classify service
type Classify struct {
	Category categories.ICategories
	Supplier suppliers.ISuppliers
}

// CreateCategory implements IClassify.
func (c *Classify) CreateCategory(ctx context.Context, ctg entity.Categories) error {
	return c.Category.Insert(ctx, ctg)
}

// CreateSuppliers implements IClassify.
func (c *Classify) CreateSuppliers(ctx context.Context, supplierReq dtos.Supplier) error {
	tx, err := db.NewTransaction(ctx)
	if err != nil {
		return err
	}
	var (
		supplier = entity.Suppliers{
			Name:  supplierReq.Name,
			Email: supplierReq.Email,
		}
		supplierRepo = suppliers.New(tx)
	)
	if err := supplierRepo.Insert(ctx, supplier); err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return err
	}
	return tx.Commit(ctx)
}

// GetCategoriesLimit implements IClassify.
func (c *Classify) GetCategoriesLimit(ctx context.Context, limit string) ([]entity.Categories, error) {
	return c.Category.GetLimit(ctx, limit)
}

// GetSuppliersLimit implements IClassify.
func (c *Classify) GetSuppliersLimit(ctx context.Context, limit int) ([]entity.Suppliers, error) {
	return c.Supplier.GetLimit(ctx, limit)
}

// DelCategoryByID implements IProductService.
func (c *Classify) DelCategoryByID(ctx context.Context, categoryID int64) error {
	return c.Category.DeleteByID(ctx, categoryID)
}

// UpdateCategoryInfo implements IProductService.
func (c *Classify) UpdateCategoryInfo(ctx context.Context, category dtos.UpdateCategories) error {

	_category := entity.Categories{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return c.Category.Update(ctx, _category)
}
