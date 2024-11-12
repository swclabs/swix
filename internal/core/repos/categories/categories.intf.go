package categories

import (
	"context"

	"github.com/swclabs/swipex/internal/core/domain/entity"
)

// ICategories defines methods to interact with category (Categories) data.
type ICategories interface {
	// Insert adds a new category to the database.
	// ctx is the context to manage the request's lifecycle.
	// ctg is a pointer to the Categories object to be added.
	// Returns an error if any issues occur during the insertion process.
	Insert(ctx context.Context, ctg entity.Category) error

	// GetLimit retrieves a list of categories with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of categories to retrieve.
	// Returns a slice of Categories objects and an error if any issues occur during the retrieval process.
	GetLimit(ctx context.Context, limit string) ([]entity.Category, error)

	// GetByID retrieves a category by its ID.
	// ctx is the context to manage the request's lifecycle.
	// ID is the ID of the category to retrieve.
	// Returns a pointer to the Categories object and an error if any issues occur during the retrieval process.
	GetByID(ctx context.Context, ID int64) (*entity.Category, error)

	// DeleteByID deletes a category by its ID.
	// ctx is the context to manage the request's lifecycle.
	// ID is the ID of the category to delete.
	// Returns an error if any issues occur during the deletion process.
	DeleteByID(ctx context.Context, ID int64) error

	// Update updates a category's information in the database.
	// ctx is the context to manage the request's lifecycle.
	// category contains the updated category details.
	// Returns an error if any issues occur during the update process.
	Update(ctx context.Context, category entity.Category) error
}
