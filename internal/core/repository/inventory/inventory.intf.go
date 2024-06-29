package inventory

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

type IInventoryRepository interface {
	InsertProduct(ctx context.Context, product domain.InventoryStruct) error
	GetProducts(ctx context.Context, productID, ram, ssd, color string) (*domain.Inventory, error)
	GetById(ctx context.Context, inventoryId int64) (*domain.Inventory, error)
}
