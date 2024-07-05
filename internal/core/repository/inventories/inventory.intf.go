package inventories

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

type IInventoryRepository interface {
	InsertProduct(ctx context.Context, product domain.InventoryStruct) error
	FindDevice(ctx context.Context, deviceSpecs domain.InventoryDeviveSpecs) (*domain.Inventories, error)
	GetById(ctx context.Context, inventoryId int64) (*domain.Inventories, error)
	GetByProductId(ctx context.Context, productId int64) ([]domain.Inventories, error)
	GetLimit(ctx context.Context, limit int, offset int) ([]domain.Inventories, error)
}
