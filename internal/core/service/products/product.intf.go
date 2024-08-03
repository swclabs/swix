// Package products implements products
package products

import (
	"context"
	"mime/multipart"
	"swclabs/swipecore/internal/core/domain/dtos"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/internal/core/domain/enum"
)

// IProductService : Module for Product interactions.
// Actor: Admin & Customer (Users)
type IProductService interface {

	// GetProductsLimit retrieves a list of products with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of products to retrieve.
	// Returns a slice of ProductResponse objects and an error if any issues occur during the retrieval process.
	GetProductsLimit(ctx context.Context, limit int) ([]dtos.ProductResponse, error)

	// UploadProductImage uploads images for a product.
	// ctx is the context to manage the request's lifecycle.
	// Id is the ID of the product.
	// fileHeader contains the file headers of the images to be uploaded.
	// Returns an error if any issues occur during the upload process.
	UploadProductImage(ctx context.Context, ID int, fileHeader []*multipart.FileHeader) error

	// CreateProduct adds a new product to the database.
	// ctx is the context to manage the request's lifecycle.
	// products contains the product details to be added.
	// Returns the ID of the newly inserted product and an error if any issues occur during the insertion process.
	CreateProduct(ctx context.Context, products dtos.Product) (int64, error)

	// InsertIntoInventory adds a product to the Inventories.
	// ctx is the context to manage the request's lifecycle.
	// product contains the inventories product details to be added.
	// Returns an error if any issues occur during the insertion process.
	InsertIntoInventory(ctx context.Context, product dtos.Inventory) error

	// DeleteProductByID deletes a product from the database.
	// ctx is the context to manage the request's lifecycle.
	// productID is the ID of the product to be deleted.
	// Returns an error if any issues occur during the deletion process.
	DeleteProductByID(ctx context.Context, productID int64) error

	// UpdateProductInfo updates a product's information in the database.
	// ctx is the context to manage the request's lifecycle.
	// product contains the updated product details.
	// Returns an error if any issues occur during the update process.
	UpdateProductInfo(ctx context.Context, product dtos.UpdateProductInfo) error

	// GetInventory retrieves a list of inventories for a product.
	// ctx is the context to manage the request's lifecycle.
	// productID is the ID of the product to retrieve inventories for.
	// Returns a slice of Inventories objects and an error if any issues occur during the retrieval process.
	GetInventory(ctx context.Context, productID int64) ([]entity.Inventories, error)

	// Search retrieves a list of products based on a search keyword.
	// ctx is the context to manage the request's lifecycle.
	// keyword is the search keyword.
	// Returns a slice of ProductResponse objects and an error if any issues occur during the retrieval process.
	Search(ctx context.Context, keyword string) ([]dtos.ProductResponse, error)

	// GetAllStock retrieves a list of all stock.
	// ctx is the context to manage the request's lifecycle.
	// page is the page number.
	// limit is the maximum number of stock to retrieve.
	// Returns a pointer to the StockInInventory object and an error if any issues occur during the retrieval process.
	GetAllStock(ctx context.Context, page int, limit int) (*dtos.StockInInventory, error)

	// DeleteInventoryByID deletes an inventory by its ID.
	// ctx is the context to manage the request's lifecycle.
	// inventoryID is the ID of the inventory to be deleted.
	// Returns an error if any issues occur during the deletion process.
	DeleteInventoryByID(ctx context.Context, inventoryID int64) error

	// UploadStockImage uploads images for an inventory.
	// ctx is the context to manage the request's lifecycle.
	// ID is the ID of the inventory.
	// fileHeader contains the file headers of the images to be uploaded.
	UploadStockImage(ctx context.Context, ID int, fileHeader []*multipart.FileHeader) error

	// UpdateInventory updates an inventory.
	// ctx is the context to manage the request's lifecycle.
	// inventory contains the updated inventory details.
	// Returns an error if any issues occur during the update process.
	UpdateInventory(ctx context.Context, inventory dtos.UpdateInventory) error

	// ProductDetailOf retrieves the details of a product.
	// ctx is the context to manage the request's lifecycle.
	// productID is the ID of the product to retrieve details for.
	// Returns a pointer to the ProductDetail object and an error if any issues occur during the retrieval
	ProductDetailOf(ctx context.Context, productID int64) (*dtos.ProductDetail, error)

	// GetInventoryByID retrieves an inventory by its ID.
	// ctx is the context to manage the request's lifecycle.
	// inventoryID is the ID of the inventory to retrieve.
	// Returns a pointer to the Inventory object and an error if any issues occur during the retrieval process.
	GetInventoryByID(ctx context.Context, inventoryID int64) (*dtos.Inventory, error)

	// ViewDataOf retrieves the data of a product.
	// ctx is the context to manage the request's lifecycle.
	// types is the category of the product.
	// Returns a slice of ProductView objects and an error if any issues occur during the retrieval process.
	ViewDataOf(ctx context.Context, types enum.Category, offset int) ([]dtos.ProductView, error)
}
