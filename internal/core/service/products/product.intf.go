// Package products implements products
package products

import (
	"context"
	"mime/multipart"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/domain/enum"
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
	// ID is the ID of the product.
	// fileHeader contains the file headers of the images to be uploaded.
	// Returns an error if any issues occur during the upload process.
	UploadProductImage(ctx context.Context, ID int, fileHeader []*multipart.FileHeader) error

	// CreateProduct adds a new product to the database.
	// ctx is the context to manage the request's lifecycle.
	// products contains the product details to be added.
	// Returns the ID of the newly inserted product and an error if any issues occur during the insertion process.
	CreateProduct(ctx context.Context, products dtos.Product) (int64, error)

	// InsertInv adds a product to the Inventories.
	// ctx is the context to manage the request's lifecycle.
	// product contains the inventories product details to be added.
	// Returns an error if any issues occur during the insertion process.
	InsertInv(ctx context.Context, product dtos.Inventory[interface{}]) error

	// DelProductByID deletes a product from the database.
	// ctx is the context to manage the request's lifecycle.
	// productID is the ID of the product to be deleted.
	// Returns an error if any issues occur during the deletion process.
	DelProductByID(ctx context.Context, productID int64) error

	// UpdateProductInfo updates a product's information in the database.
	// ctx is the context to manage the request's lifecycle.
	// product contains the updated product details.
	// Returns an error if any issues occur during the update process.
	UpdateProductInfo(ctx context.Context, product dtos.UpdateProductInfo) error

	// GetInv retrieves a list of inventories for a product.
	// ctx is the context to manage the request's lifecycle.
	// productID is the ID of the product to retrieve inventories for.
	// Returns a slice of Inventories objects and an error if any issues occur during the retrieval process.
	GetInv(ctx context.Context, productID int64) ([]entity.Inventories, error)

	// Search retrieves a list of products based on a search keyword.
	// ctx is the context to manage the request's lifecycle.
	// keyword is the search keyword.
	// Returns a slice of ProductResponse objects and an error if any issues occur during the retrieval process.
	Search(ctx context.Context, keyword string) ([]dtos.ProductResponse, error)

	// GetAllInv retrieves a list of all stock.
	// ctx is the context to manage the request's lifecycle.
	// page is the page number.
	// limit is the maximum number of stock to retrieve.
	// Returns a pointer to the InvStock object and an error if any issues occur during the retrieval process.
	GetAllInv(ctx context.Context, page int, limit int) (*dtos.InvStock[interface{}], error)

	// DeleteInvByID deletes an inventory by its ID.
	// ctx is the context to manage the request's lifecycle.
	// inventoryID is the ID of the inventory to be deleted.
	// Returns an error if any issues occur during the deletion process.
	DeleteInvByID(ctx context.Context, inventoryID int64) error

	// UploadInvImage uploads images for an inventory.
	// ctx is the context to manage the request's lifecycle.
	// ID is the ID of the inventory.
	// fileHeader contains the file headers of the images to be uploaded.
	UploadInvImage(ctx context.Context, ID int, fileHeader []*multipart.FileHeader) error

	// UpdateInv updates an inventory.
	// ctx is the context to manage the request's lifecycle.
	// inventory contains the updated inventory details.
	// Returns an error if any issues occur during the update process.
	UpdateInv(ctx context.Context, inventory dtos.InvUpdate) error

	// ProductDetail retrieves the details of a product.
	// ctx is the context to manage the request's lifecycle.
	// productID is the ID of the product to retrieve details for.
	// Returns a pointer to the ProductDetail object and an error if any issues occur during the retrieval
	ProductDetail(ctx context.Context, productID int64) (*dtos.ProductDetail[interface{}], error)

	// GetInvByID retrieves an inventory by its ID.
	// ctx is the context to manage the request's lifecycle.
	// inventoryID is the ID of the inventory to retrieve.
	// Returns a pointer to the Inventory object and an error if any issues occur during the retrieval process.
	GetInvByID(ctx context.Context, inventoryID int64) (*dtos.Inventory[interface{}], error)

	// ViewDataOf retrieves the data of a product.
	// ctx is the context to manage the request's lifecycle.
	// types is the category of the product.
	// Returns a slice of ProductView objects and an error if any issues occur during the retrieval process.
	ViewDataOf(ctx context.Context, types enum.Category, offset int) ([]dtos.ProductView, error)

	// InsertSpecStorage inserts specifications for an inventory.
	// ctx is the context to manage the request's lifecycle.
	// specification contains the specifications to be added.
	// Returns an error if any issues occur during the insertion process.
	InsertSpecStorage(ctx context.Context, specification dtos.Storage) error

	InsertSpecWireless(ctx context.Context, specification dtos.Wireless) error
}
