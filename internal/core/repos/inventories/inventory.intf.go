package inventories

import (
	"context"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/internal/core/domain/model"
)

// IInventories represents the interface for Inventory repos.
type IInventories interface {
	// InsertProduct inserts a product to the inventory.
	InsertProduct(ctx context.Context, product entity.Inventory) (int64, error)

	// GetByID gets an inventory by its ID.
	GetByID(ctx context.Context, inventoryID int64) (*entity.Inventory, error)

	// GetByProductID gets inventories by product ID.
	GetByProductID(ctx context.Context, productID int64) ([]entity.Inventory, error)

	// GetLimit gets inventories with limit and offset.
	GetLimit(ctx context.Context, limit int, offset int) ([]entity.Inventory, error)

	// DeleteByID deletes an inventory by its ID.
	DeleteByID(ctx context.Context, inventoryID int64) error

	// UploadImage uploads an image to the inventory.
	UploadImage(ctx context.Context, ID int, url string) error

	UploadColorImage(ctx context.Context, ID int, url string) error

	// Update updates an inventory.
	Update(ctx context.Context, inventory entity.Inventory) error

	GetColor(ctx context.Context, productID int64) ([]model.ColorItem, error)

	GetByColor(ctx context.Context, productID int64, color string) ([]entity.Inventory, error)
}
