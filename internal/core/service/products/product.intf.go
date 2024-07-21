// Package products implements products
package products

import (
	"context"
	"mime/multipart"
	"swclabs/swipecore/internal/core/domain"
)

// IProductService : Module for Product interactions.
// Actor: Admin & Customer (Users)
type IProductService interface {
	// GetCategoriesLimit retrieves a list of categories with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of categories to retrieve.
	// Returns a slice of Categories objects and an error if any issues occur during the retrieval process.
	GetCategoriesLimit(ctx context.Context, limit string) ([]domain.Categories, error)

	// GetProductsLimit retrieves a list of products with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of products to retrieve.
	// Returns a slice of ProductSchema objects and an error if any issues occur during the retrieval process.
	GetProductsLimit(ctx context.Context, limit int) ([]domain.ProductSchema, error)

	// CreateCategory adds a new category to the database.
	// ctx is the context to manage the request's lifecycle.
	// ctg is a pointer to the Categories object to be added.
	// Returns an error if any issues occur during the insertion process.
	CreateCategory(ctx context.Context, ctg domain.Categories) error

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
	CreateProduct(ctx context.Context, products domain.Product) (int64, error)

	// GetSuppliersLimit retrieves a list of suppliers with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of suppliers to retrieve.
	// Returns a slice of Suppliers objects and an error if any issues occur during the retrieval process.
	GetSuppliersLimit(ctx context.Context, limit int) ([]domain.Suppliers, error)

	// CreateSuppliers adds a new supplier to the database.
	// ctx is the context to manage the request's lifecycle.
	// supplierReq contains the supplier details to be added.
	// Returns an error if any issues occur during the insertion process.
	CreateSuppliers(ctx context.Context, supplierReq domain.SupplierSchema) error

	// InsertIntoInventory adds a product to the Inventories.
	// ctx is the context to manage the request's lifecycle.
	// product contains the inventories product details to be added.
	// Returns an error if any issues occur during the insertion process.
	InsertIntoInventory(ctx context.Context, product domain.InventoryStruct) error

	// FindDeviceInInventory retrieves device details from the Inventories.
	// ctx is the context to manage the request's lifecycle.
	// deviceSpecs specify the product attributes to retrieve.
	// Returns a pointer to the InventorySchema object and an error if any issues occur during the retrieval process.
	FindDeviceInInventory(ctx context.Context, deviceSpecs domain.InventoryDeviveSpecs) (*domain.InventorySchema, error)

	// DeleteProductByID deletes a product from the database.
	// ctx is the context to manage the request's lifecycle.
	// productID is the ID of the product to be deleted.
	// Returns an error if any issues occur during the deletion process.
	DeleteProductByID(ctx context.Context, productID int64) error

	// UpdateProductInfo updates a product's information in the database.
	// ctx is the context to manage the request's lifecycle.
	// product contains the updated product details.
	// Returns an error if any issues occur during the update process.
	UpdateProductInfo(ctx context.Context, product domain.UpdateProductInfo) error

	// GetInventory retrieves a list of inventories for a product.
	// ctx is the context to manage the request's lifecycle.
	// productID is the ID of the product to retrieve inventories for.
	// Returns a slice of Inventories objects and an error if any issues occur during the retrieval process.
	GetInventory(ctx context.Context, productID int64) ([]domain.Inventories, error)

	// Search retrieves a list of products based on a search keyword.
	// ctx is the context to manage the request's lifecycle.
	// keyword is the search keyword.
	// Returns a slice of ProductSchema objects and an error if any issues occur during the retrieval process.
	Search(ctx context.Context, keyword string) ([]domain.ProductSchema, error)

	// GetAllStock retrieves a list of all stock.
	// ctx is the context to manage the request's lifecycle.
	// page is the page number.
	// limit is the maximum number of stock to retrieve.
	// Returns a pointer to the InventoryStockSchema object and an error if any issues occur during the retrieval process.
	GetAllStock(ctx context.Context, page int, limit int) (*domain.InventoryStockSchema, error)

	// DeleteInventoryByID deletes an inventory by its ID.
	// ctx is the context to manage the request's lifecycle.
	// inventoryID is the ID of the inventory to be deleted.
	// Returns an error if any issues occur during the deletion process.
	DeleteInventoryByID(ctx context.Context, inventoryID int64) error
}
