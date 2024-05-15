package domain

import "context"

type IWarehouseRepository interface {
	InsertProduct(ctx context.Context, product WarehouseReq) error
	GetProducts(ctx context.Context, productID, ram, ssd, color string) (*WarehouseRes, error)
}
