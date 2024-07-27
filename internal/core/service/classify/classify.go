// Package classify implements classify service include category and supplier
package classify

import (
	"context"
	"log"
	"swclabs/swipecore/internal/core/domain/dtos"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/internal/core/repository/addresses"
	"swclabs/swipecore/internal/core/repository/categories"
	"swclabs/swipecore/internal/core/repository/suppliers"
	"swclabs/swipecore/pkg/infra/db"

	"github.com/google/uuid"
)

// New creates a new Classify object
func New(
	category categories.ICategoriesRepository,
	supplier suppliers.ISuppliersRepository,
) IClassify {
	return &Classify{
		Category: category,
		Supplier: supplier,
	}
}

// Classify struct for classify service
type Classify struct {
	Category categories.ICategoriesRepository
	Supplier suppliers.ISuppliersRepository
}

// CreateCategory implements IClassify.
func (c *Classify) CreateCategory(ctx context.Context, ctg entity.Categories) error {
	return c.Category.Insert(ctx, ctg)
}

// CreateSuppliers implements IClassify.
func (c *Classify) CreateSuppliers(ctx context.Context, supplierReq dtos.Supplier) error {
	tx, err := db.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	var (
		supplier = entity.Suppliers{
			Name:  supplierReq.Name,
			Email: supplierReq.Email,
		}
		addr = entity.Addresses{
			City:     supplierReq.City,
			Ward:     supplierReq.Ward,
			District: supplierReq.District,
			Street:   supplierReq.Street,
		}
		supplierRepo = suppliers.New(tx)
		addressRepo  = addresses.New(tx)
	)
	if err := supplierRepo.Insert(ctx, supplier); err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return err
	}
	supp, err := supplierRepo.GetByPhone(ctx, supplierReq.Email)
	if err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return err
	}
	addr.UUID = uuid.New().String()
	if err = addressRepo.Insert(ctx, addr); err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return err
	}
	if err := supplierRepo.InsertAddress(ctx, entity.SuppliersAddress{
		SuppliersID: supp.ID,
		AddressUuiD: addr.UUID,
	}); err != nil {
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
