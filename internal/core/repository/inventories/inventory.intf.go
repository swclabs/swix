package inventories

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

type IInventoryRepository interface {
	InsertProduct(ctx context.Context, product domain.InventoryStruct) error
	FindDevice(ctx context.Context, deviceSpecs domain.InventoryDeviveSpecs) (*domain.Inventories, error)
	GetByID(ctx context.Context, inventoryID int64) (*domain.Inventories, error)
	GetByProductID(ctx context.Context, productID int64) ([]domain.Inventories, error)
	GetLimit(ctx context.Context, limit int, offset int) ([]domain.Inventories, error)
}
