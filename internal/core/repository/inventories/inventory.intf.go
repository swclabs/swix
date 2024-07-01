package inventories

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

type IInventoryRepository interface {
	InsertProduct(ctx context.Context, product domain.InventoryStruct) error
	FindDevice(ctx context.Context, productID, ram, ssd, color string) (*domain.Inventories, error)
	GetById(ctx context.Context, inventoryId int64) (*domain.Inventories, error)
	GetByProductId(ctx context.Context, productId int64) ([]domain.Inventories, error)
}
