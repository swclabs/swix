package domain

import "context"

type IWarehouseRepository interface {
	InsertProduct(ctx context.Context, product Warehouse) error
	GetProducts(ctx context.Context, productID, ram, ssd string) (*Warehouse, error)
}
