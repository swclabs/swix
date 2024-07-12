package products

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

// IProductRepository defines methods to interact with product (Products) data.
type IProductRepository interface {
	// Insert adds a new product to the database.
	// ctx is the context to manage the request's lifecycle.
	// prd is a pointer to the Products object to be added.
	// Returns the ID of the newly inserted product and an error if any issues occur during the insertion process.
	Insert(ctx context.Context, prd domain.Products) (int64, error)

	// GetLimit retrieves a list of products with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of products to retrieve.
	// Returns a slice of ProductSchema objects and an error if any issues occur during the retrieval process.
	GetLimit(ctx context.Context, limit int) ([]domain.Products, error)

	// UploadNewImage updates the image URL of a specified product.
	// ctx is the context to manage the request's lifecycle.
	// urlImg is the new image URL to be uploaded.
	// id is the ID of the product to be updated.
	// Returns an error if any issues occur during the update process.
	UploadNewImage(ctx context.Context, urlImg string, ID int) error

	GetByID(ctx context.Context, productID int64) (*domain.Products, error)
	DeleteByID(ctx context.Context, ID int64) error
	Update(ctx context.Context, product domain.Products) error
	Search(ctx context.Context, keyword string) ([]domain.Products, error)
}
