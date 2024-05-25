package domain

import "context"

type IWarehouseRepository interface {
	InsertProduct(ctx context.Context, product WarehouseStructure) error
	GetProducts(ctx context.Context, productID, ram, ssd, color string) (*Warehouse, error)
}
