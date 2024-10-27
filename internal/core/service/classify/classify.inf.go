package classify

import (
	"context"
	"swclabs/swipex/internal/core/domain/dtos"
	"swclabs/swipex/internal/core/domain/entity"
)

// IClassify : Classify utility methods for the service.
type IClassify interface {
	// CreateCategory adds a new category to the database.
	// ctx is the context to manage the request's lifecycle.
	// ctg is a pointer to the Categories object to be added.
	// Returns an error if any issues occur during the insertion process.
	CreateCategory(ctx context.Context, ctg entity.Category) error

	// GetCategoriesLimit retrieves a list of categories with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of categories to retrieve.
	// Returns a slice of Categories objects and an error if any issues occur during the retrieval process.
	GetCategoriesLimit(ctx context.Context, limit string) ([]entity.Category, error)

	// GetSuppliersLimit retrieves a list of suppliers with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of suppliers to retrieve.
	// Returns a slice of Suppliers objects and an error if any issues occur during the retrieval process.
	GetSuppliersLimit(ctx context.Context, limit int) ([]entity.Supplier, error)

	// CreateSuppliers adds a new supplier to the database.
	// ctx is the context to manage the request's lifecycle.
	// supplierReq contains the supplier details to be added.
	// Returns an error if any issues occur during the insertion process.
	CreateSuppliers(ctx context.Context, supplierReq dtos.Supplier) error

	// DeleteCategoryByID deletes a category from the database.
	// ctx is the context to manage the request's lifecycle.
	// categoryID is the ID of the category to be deleted.
	// Returns an error if any issues occur during the deletion process.
	DelCategoryByID(ctx context.Context, categoryID int64) error

	// UpdateCategoryInfo updates a category's information in the database.
	// ctx is the context to manage the request's lifecycle.
	// category contains the updated category details.
	// Returns an error if any issues occur during the update process.
	UpdateCategoryInfo(ctx context.Context, ctg dtos.UpdateCategories) error
}
