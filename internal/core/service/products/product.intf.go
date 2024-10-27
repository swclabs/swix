// Package products implements products
package products

import (
	"context"
	"mime/multipart"
	"swclabs/swipex/internal/core/domain/dtos"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/internal/core/domain/enum"
)

// IProducts : Module for Product interactions.
// Actor: Admin & Customer (Users)
type IProducts interface {
	// GetProducts retrieves a list of products with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of products to retrieve.
	// Returns a slice of ProductResponse objects and an error if any issues occur during the retrieval process.
	GetProducts(ctx context.Context, limit int) ([]dtos.ProductResponse, error)

	// UploadProductImage uploads images for a product.
	// ctx is the context to manage the request's lifecycle.
	// ID is the ID of the product.
	// fileHeader contains the file headers of the images to be uploaded.
	// Returns an error if any issues occur during the upload process.
	UploadProductImage(ctx context.Context, ID int, fileHeader []*multipart.FileHeader) error

	// UploadProductShopImage uploads shop images for a product.
	// ctx is the context to manage the request's lifecycle.
	// ID is the ID of the product.
	// fileHeader contains the file headers of the images to be uploaded.
	// Returns an error if any issues occur during the upload process.
	UploadProductShopImage(ctx context.Context, ID int, fileHeader []*multipart.FileHeader) error

	// CreateProduct adds a new product to the database.
	// ctx is the context to manage the request's lifecycle.
	// products contains the product details to be added.
	// Returns the ID of the newly inserted product and an error if any issues occur during the insertion process.
	CreateProduct(ctx context.Context, products dtos.Product) (int64, error)

	// InsertItem adds a product to the Inventories.
	// ctx is the context to manage the request's lifecycle.
	// product contains the inventories product details to be added.
	// Returns an error if any issues occur during the insertion process.
	InsertItem(ctx context.Context, product dtos.Inventory) error

	// DelProduct deletes a product from the database.
	// ctx is the context to manage the request's lifecycle.
	// productID is the ID of the product to be deleted.
	// Returns an error if any issues occur during the deletion process.
	DelProduct(ctx context.Context, productID int64) error

	// UpdateProductInfo updates a product's information in the database.
	// ctx is the context to manage the request's lifecycle.
	// product contains the updated product details.
	// Returns an error if any issues occur during the update process.
	UpdateProductInfo(ctx context.Context, product dtos.UpdateProductInfo) error

	// GetItems retrieves a list of inventories for a product.
	// ctx is the context to manage the request's lifecycle.
	// productID is the ID of the product to retrieve inventories for.
	// Returns a slice of Inventories objects and an error if any issues occur during the retrieval process.
	GetItems(ctx context.Context, productID int64) ([]entity.Inventory, error)

	// Search retrieves a list of products based on a search keyword.
	// ctx is the context to manage the request's lifecycle.
	// keyword is the search keyword.
	// Returns a slice of ProductResponse objects and an error if any issues occur during the retrieval process.
	Search(ctx context.Context, keyword string) ([]dtos.ProductResponse, error)

	// SearchDetails retrieves a list of products based on a search keyword.
	// ctx is the context to manage the request's lifecycle.
	// keyword is the search keyword.
	// Returns a slice of ProductDetail objects and an error if any issues occur during the retrieval process.
	SearchDetails(ctx context.Context, keyword string) ([]dtos.ProductDetail, error)

	// GetInvItems retrieves a list of all stock.
	// ctx is the context to manage the request's lifecycle.
	// page is the page number.
	// limit is the maximum number of stock to retrieve.
	// Returns a pointer to the InvStock object and an error if any issues occur during the retrieval process.
	GetInvItems(ctx context.Context, page int, limit int) (*dtos.InventoryItems, error)

	// DeleteItem deletes an inventory by its ID.
	// ctx is the context to manage the request's lifecycle.
	// inventoryID is the ID of the inventory to be deleted.
	// Returns an error if any issues occur during the deletion process.
	DeleteItem(ctx context.Context, inventoryID int64) error

	// UploadItemImage uploads images for an inventory.
	// ctx is the context to manage the request's lifecycle.
	// ID is the ID of the inventory.
	// fileHeader contains the file headers of the images to be uploaded.
	UploadItemImage(ctx context.Context, ID int, fileHeader []*multipart.FileHeader) error

	// UploadItemColorImage uploads color images for an inventory.
	// ctx is the context to manage the request's lifecycle.
	// ID is the ID of the inventory.
	// fileHeader contains the file headers of the images to be uploaded.
	UploadItemColorImage(ctx context.Context, ID int, fileHeader []*multipart.FileHeader) error

	// UpdateItem updates an inventory.
	// ctx is the context to manage the request's lifecycle.
	// inventory contains the updated inventory details.
	// Returns an error if any issues occur during the update process.
	UpdateItem(ctx context.Context, inventory dtos.InvUpdate) error

	// Detail retrieves the details of a product.
	// ctx is the context to manage the request's lifecycle.
	// productID is the ID of the product to retrieve details for.
	// Returns a pointer to the Detail object and an error if any issues occur during the retrieval
	Detail(ctx context.Context, productID int64) (*dtos.ProductDetail, error)

	// GetItem retrieves an inventory by its ID.
	// ctx is the context to manage the request's lifecycle.
	// inventoryID is the ID of the inventory to retrieve.
	// Returns a pointer to the Inventory object and an error if any issues occur during the retrieval process.
	GetItem(ctx context.Context, inventoryID int64) (*dtos.Inventory, error)

	// ProductOf retrieves the data of a product.
	// ctx is the context to manage the request's lifecycle.
	// types is the category of the product.
	// Returns a slice of ProductView objects and an error if any issues occur during the retrieval process.
	ProductOf(ctx context.Context, types enum.Category, offset int) ([]dtos.ProductDTO, error)

	// Rating updates the rating of a product.
	// ctx is the context to manage the request's lifecycle.
	// productID is the ID of the product to update the rating for.
	// rating is the new rating of the product.
	// Returns an error if any issues occur during the update process.
	Rating(ctx context.Context, userID, productID int64, rating float64) error
}
