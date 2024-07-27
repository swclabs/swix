package inventories

import (
	"context"
	"swclabs/swipecore/internal/core/domain/dtos"
	"swclabs/swipecore/internal/core/domain/entity"
)

// IInventoryRepository represents the interface for Inventory repository.
type IInventoryRepository interface {
	// InsertProduct inserts a product to the inventory.
	InsertProduct(ctx context.Context, product entity.Inventories) error

	// FindDevice finds a device in the inventory.
	FindDevice(ctx context.Context, device dtos.InventoryDeviceSpecs) (*entity.Inventories, error)

	// GetByID gets an inventory by its ID.
	GetByID(ctx context.Context, inventoryID int64) (*entity.Inventories, error)

	// GetByProductID gets inventories by product ID.
	GetByProductID(ctx context.Context, productID int64) ([]entity.Inventories, error)

	// GetLimit gets inventories with limit and offset.
	GetLimit(ctx context.Context, limit int, offset int) ([]entity.Inventories, error)

	// DeleteByID deletes an inventory by its ID.
	DeleteByID(ctx context.Context, inventoryID int64) error

	// UploadImage uploads an image to the inventory.
	UploadImage(ctx context.Context, ID int, url string) error

	// Update updates an inventory.
	Update(ctx context.Context, inventory entity.Inventories) error
}
