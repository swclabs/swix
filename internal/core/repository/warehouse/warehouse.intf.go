package warehouse

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

type IWarehouseRepository interface {
	InsertProduct(ctx context.Context, product domain.WarehouseStruct) error
	GetProducts(ctx context.Context, productID, ram, ssd, color string) (*domain.Warehouse, error)
}
