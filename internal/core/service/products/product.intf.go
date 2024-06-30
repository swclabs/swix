package products

import (
	"context"
	"mime/multipart"
	"swclabs/swipecore/internal/core/domain"
)

// IProductService : Module for Product interactions.
// Actor: Admin & Customer (User)
type IProductService interface {
	// GetCategoriesLimit retrieves a list of categories with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of categories to retrieve.
	// Returns a slice of Categories objects and an error if any issues occur during the retrieval process.
	GetCategoriesLimit(ctx context.Context, limit string) ([]domain.Categories, error)

	// GetProductsLimit retrieves a list of products with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of products to retrieve.
	// Returns a slice of ProductRes objects and an error if any issues occur during the retrieval process.
	GetProductsLimit(ctx context.Context, limit int) ([]domain.ProductRes, error)

	// InsertCategory adds a new category to the database.
	// ctx is the context to manage the request's lifecycle.
	// ctg is a pointer to the Categories object to be added.
	// Returns an error if any issues occur during the insertion process.
	InsertCategory(ctx context.Context, ctg domain.Categories) error

	// UploadProductImage uploads images for a product.
	// ctx is the context to manage the request's lifecycle.
	// Id is the ID of the product.
	// fileHeader contains the file headers of the images to be uploaded.
	// Returns an error if any issues occur during the upload process.
	UploadProductImage(ctx context.Context, Id int, fileHeader []*multipart.FileHeader) error

	// UploadProduct adds a new product to the database.
	// ctx is the context to manage the request's lifecycle.
	// products contains the product details to be added.
	// Returns the ID of the newly inserted product and an error if any issues occur during the insertion process.
	UploadProduct(ctx context.Context, products domain.ProductReq) (int64, error)

	// GetSuppliersLimit retrieves a list of suppliers with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of suppliers to retrieve.
	// Returns a slice of Suppliers objects and an error if any issues occur during the retrieval process.
	GetSuppliersLimit(ctx context.Context, limit int) ([]domain.Suppliers, error)

	// InsertSuppliers adds a new supplier to the database.
	// ctx is the context to manage the request's lifecycle.
	// supplierReq contains the supplier details to be added.
	// Returns an error if any issues occur during the insertion process.
	InsertSuppliers(ctx context.Context, supplierReq domain.SuppliersReq) error

	// InsertIntoInventory adds a product to the Inventory.
	// ctx is the context to manage the request's lifecycle.
	// product contains the inventory product details to be added.
	// Returns an error if any issues occur during the insertion process.
	InsertIntoInventory(ctx context.Context, product domain.InventoryStruct) error

	// GetProductsInInventory retrieves product details from the Inventory.
	// ctx is the context to manage the request's lifecycle.
	// productID, ram, ssd, and color specify the product attributes to retrieve.
	// Returns a pointer to the InventorySchema object and an error if any issues occur during the retrieval process.
	GetProductsInInventory(ctx context.Context, productID, ram, ssd, color string) (*domain.InventorySchema, error)

	DeleteProductById(ctx context.Context, productId int64) error
	UpdateProductInfor(ctx context.Context, product domain.UpdateProductInfoReq) error
}
