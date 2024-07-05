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
	UploadProductImage(ctx context.Context, Id int, fileHeader []*multipart.FileHeader) error

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

	DeleteProductById(ctx context.Context, productId int64) error
	UpdateProductInfo(ctx context.Context, product domain.UpdateProductInfo) error
	GetInventory(ctx context.Context, productId int64) ([]domain.Inventories, error)
	Search(ctx context.Context, keyword string) ([]domain.ProductSchema, error)
	GetAllStock(ctx context.Context, page int, limit int) (*domain.InventoryStockSchema, error)
}
