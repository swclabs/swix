package categories

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

// ICategoriesRepository defines methods to interact with category (Categories) data.
type ICategoriesRepository interface {
	// Insert adds a new category to the database.
	// ctx is the context to manage the request's lifecycle.
	// ctg is a pointer to the Categories object to be added.
	// Returns an error if any issues occur during the insertion process.
	Insert(ctx context.Context, ctg *domain.Categories) error

	// GetLimit retrieves a list of categories with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of categories to retrieve.
	// Returns a slice of Categories objects and an error if any issues occur during the retrieval process.
	GetLimit(ctx context.Context, limit string) ([]domain.Categories, error)
}
