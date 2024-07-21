package inventories

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

// IInventoryRepository represents the interface for Inventory repository.
type IInventoryRepository interface {
	// InsertProduct inserts a product to the inventory.
	InsertProduct(ctx context.Context, product domain.InventoryStruct) error

	// FindDevice finds a device in the inventory.
	FindDevice(ctx context.Context, deviceSpecs domain.InventoryDeviveSpecs) (*domain.Inventories, error)

	// GetByID gets an inventory by its ID.
	GetByID(ctx context.Context, inventoryID int64) (*domain.Inventories, error)

	// GetByProductID gets inventories by product ID.
	GetByProductID(ctx context.Context, productID int64) ([]domain.Inventories, error)

	// GetLimit gets inventories with limit and offset.
	GetLimit(ctx context.Context, limit int, offset int) ([]domain.Inventories, error)

	// DeleteByID deletes an inventory by its ID.
	DeleteByID(ctx context.Context, inventoryID int64) error
}
