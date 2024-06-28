package warehouse

import (
	"context"
	"encoding/json"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"

	"github.com/jackc/pgx/v5"
)

type Warehouse struct {
	db db.IDatabase
}

var _ IWarehouseRepository = (*Warehouse)(nil)

func New(conn db.IDatabase) IWarehouseRepository {
	return &Warehouse{
		db: conn,
	}
}

// GetById implements IWarehouseRepository.
func (w *Warehouse) GetById(ctx context.Context, warehouseId int64) (*domain.Warehouse, error) {
	rows, err := w.db.Query(ctx, getById, warehouseId)
	if err != nil {
		return nil, err
	}
	warehouse, err := db.CollectOneRow[domain.Warehouse](rows)
	if err != nil {
		return nil, err
	}
	return &warehouse, nil
}

// GetProducts implements domain.IWarehouseRepository.
func (w *Warehouse) GetProducts(
	ctx context.Context, productID, ram, ssd, color string) (*domain.Warehouse, error) {
	rows, err := w.db.Query(ctx, getAvailableProducts, productID, ram, ssd, color)
	if err != nil {
		return nil, err
	}
	warehouse, err := pgx.CollectOneRow[domain.Warehouse](rows, pgx.RowToStructByName[domain.Warehouse])
	if err != nil {
		return nil, err
	}
	return &warehouse, nil
}

// InsertProduct implements domain.IWarehouseRepository.
func (w *Warehouse) InsertProduct(
	ctx context.Context, product domain.WarehouseStruct) error {
	specsjson, _ := json.Marshal(product.Specs)
	return w.db.SafeWrite(ctx, insertIntoWarehouse,
		product.ProductID, product.Model, product.Price,
		string(specsjson), product.Available, product.CurrencyCode,
	)
}
