package repository

import (
	"context"
	"encoding/json"
	"log"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"
	"swclabs/swipecore/pkg/db/queries"

	"gorm.io/gorm"
)

type Warehouse struct {
	conn *gorm.DB
}

var _ domain.IWarehouseRepository = (*Warehouse)(nil)

func NewWarehouse() domain.IWarehouseRepository {
	_conn, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	return &Warehouse{
		conn: _conn,
	}
}

// GetProducts implements domain.IWarehouseRepository.
func (w *Warehouse) GetProducts(ctx context.Context, productID, ram, ssd, color string) (*domain.WarehouseRes, error) {
	var warehouse domain.Warehouse
	if err := w.conn.
		WithContext(ctx).
		Raw(queries.GetAvailableProducts, productID, ram, ssd, color).
		Scan(&warehouse).Error; err != nil {
		return nil, err
	}
	var warehouseRes = domain.WarehouseRes{
		Id: warehouse.Id,
		WarehouseReq: domain.WarehouseReq{
			ProductID: warehouse.Id,
			Price:     warehouse.Price,
			Model:     warehouse.Model,
			Available: warehouse.Available,
		},
	}
	if err := json.Unmarshal([]byte(warehouse.Specs), &warehouseRes.Specs); err != nil {
		return &warehouseRes, nil // don't find anything, just return empty object
	}
	if warehouseRes.Available == "" {
		warehouseRes.Available = "0"
		return &warehouseRes, nil
	}
	return &warehouseRes, nil
}

// InsertProduct implements domain.IWarehouseRepository.
func (w *Warehouse) InsertProduct(ctx context.Context, product domain.WarehouseReq) error {
	specsjson, _ := json.Marshal(product.Specs)
	return db.SafeWriteQuery(
		ctx,
		w.conn,
		queries.InsertIntoWarehouse,
		product.ProductID, product.Model, product.Price, string(specsjson), product.Available,
	)
}
