package warehouse

import (
	"context"
	"encoding/json"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"

	"github.com/jackc/pgx/v5"
)

type Warehouse struct {
	conn *pgx.Conn
}

var _ IWarehouseRepository = (*Warehouse)(nil)

func New(conn *pgx.Conn) *Warehouse {
	return &Warehouse{
		conn: conn,
	}
}

// GetProducts implements domain.IWarehouseRepository.
func (w *Warehouse) GetProducts(
	ctx context.Context, productID, ram, ssd, color string) (*domain.Warehouse, error) {
	rows, err := w.conn.Query(ctx, GetAvailableProducts, productID, ram, ssd, color)
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
	return db.SafePgxWriteQuery(ctx, w.conn, InsertIntoWarehouse,
		product.ProductID, product.Model, product.Price,
		string(specsjson), product.Available, product.CurrencyCode,
	)
}
