package inventory

import (
	"context"
	"encoding/json"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"
)

type Inventory struct {
	db db.IDatabase
}

var _ IInventoryRepository = (*Inventory)(nil)

func New(conn db.IDatabase) IInventoryRepository {
	return useCache(&Inventory{
		db: conn,
	})
}

// GetById implements IInventoryRepository.
func (w *Inventory) GetById(ctx context.Context, inventoryId int64) (*domain.Inventory, error) {
	rows, err := w.db.Query(ctx, getById, inventoryId)
	if err != nil {
		return nil, err
	}
	inventory, err := db.CollectOneRow[domain.Inventory](rows)
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

// GetProducts implements domain.IInventoryRepository.
func (w *Inventory) GetProducts(
	ctx context.Context, productID, ram, ssd, color string) (*domain.Inventory, error) {
	rows, err := w.db.Query(ctx, getAvailableProducts, productID, ram, ssd, color)
	if err != nil {
		return nil, err
	}
	inventory, err := db.CollectOneRow[domain.Inventory](rows)
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

// InsertProduct implements domain.IInventoryRepository.
func (w *Inventory) InsertProduct(
	ctx context.Context, product domain.InventoryStruct) error {
	specsjson, _ := json.Marshal(product.Specs)
	return w.db.SafeWrite(ctx, insertIntoInventory,
		product.ProductID, product.Model, product.Price,
		string(specsjson), product.Available, product.CurrencyCode,
	)
}
