package warehouse

import (
	"context"
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"
)

type Warehouse struct {
	conn *gorm.DB
}

var _ IWarehouseRepository = (*Warehouse)(nil)

func New() IWarehouseRepository {
	_conn, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	return &Warehouse{
		conn: _conn,
	}
}

// GetProducts implements domain.IWarehouseRepository.
func (w *Warehouse) GetProducts(
	ctx context.Context, productID, ram, ssd, color string) (*domain.Warehouse, error) {
	var warehouse domain.Warehouse
	if err := w.conn.
		WithContext(ctx).
		Raw(GetAvailableProducts, productID, ram, ssd, color).
		Scan(&warehouse).Error; err != nil {
		return nil, err
	}
	return &warehouse, nil
}

// InsertProduct implements domain.IWarehouseRepository.
func (w *Warehouse) InsertProduct(ctx context.Context, product domain.WarehouseStructure) error {
	specsjson, _ := json.Marshal(product.Specs)
	return db.SafeWriteQuery(
		ctx,
		w.conn,
		InsertIntoWarehouse,
		product.ProductID, product.Model, product.Price, string(specsjson), product.Available,
	)
}
