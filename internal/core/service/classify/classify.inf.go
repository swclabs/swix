package classify

import (
	"context"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/domain/entity"
)

// IClassify : Classify utility methods for the service.
type IClassify interface {
	// CreateCategory adds a new category to the database.
	// ctx is the context to manage the request's lifecycle.
	// ctg is a pointer to the Categories object to be added.
	// Returns an error if any issues occur during the insertion process.
	CreateCategory(ctx context.Context, ctg entity.Categories) error

	// GetCategoriesLimit retrieves a list of categories with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of categories to retrieve.
	// Returns a slice of Categories objects and an error if any issues occur during the retrieval process.
	GetCategoriesLimit(ctx context.Context, limit string) ([]entity.Categories, error)

	// GetSuppliersLimit retrieves a list of suppliers with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of suppliers to retrieve.
	// Returns a slice of Suppliers objects and an error if any issues occur during the retrieval process.
	GetSuppliersLimit(ctx context.Context, limit int) ([]entity.Suppliers, error)

	// CreateSuppliers adds a new supplier to the database.
	// ctx is the context to manage the request's lifecycle.
	// supplierReq contains the supplier details to be added.
	// Returns an error if any issues occur during the insertion process.
	CreateSuppliers(ctx context.Context, supplierReq dtos.Supplier) error
}
