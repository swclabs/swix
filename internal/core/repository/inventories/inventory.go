package inventories

import (
	"context"
	"encoding/json"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/errors"
	"swclabs/swipecore/pkg/infra/db"
)

var _ IInventoryRepository = (*Inventory)(nil)

// New creates a new Inventory object.
func New(conn db.IDatabase) IInventoryRepository {
	return useCache(&Inventory{
		db: conn,
	})
}

// Inventory represents the Inventory object.
type Inventory struct {
	db db.IDatabase
}

// UploadImage implements IInventoryRepository.
func (w *Inventory) UploadImage(ctx context.Context, ID int, url string) error {
	return w.db.SafeWrite(ctx, uploadInventoryImage, url, ID)
}

// DeleteByID implements IInventoryRepository.
func (w *Inventory) DeleteByID(ctx context.Context, inventoryID int64) error {
	return w.db.SafeWrite(ctx, deleteInventorybyID, inventoryID)
}

// GetLimit implements IInventoryRepository.
func (w *Inventory) GetLimit(ctx context.Context, limit int, offset int) ([]domain.Inventories, error) {
	if offset < 1 {
		offset = 1
	}
	rows, err := w.db.Query(ctx, getProductsLimit, limit, (offset-1)*limit)
	if err != nil {
		return nil, err
	}
	return db.CollectRows[domain.Inventories](rows)
}

// GetByProductID implements IInventoryRepository.
func (w *Inventory) GetByProductID(ctx context.Context, productID int64) ([]domain.Inventories, error) {
	rows, err := w.db.Query(ctx, getByProductID, productID)
	if err != nil {
		return nil, errors.Repository("500", err)
	}
	inventories, err := db.CollectRows[domain.Inventories](rows)
	if err != nil {
		return nil, errors.Repository("500", err)
	}
	return inventories, nil
}

// GetByID implements IInventoryRepository.
func (w *Inventory) GetByID(ctx context.Context, inventoryID int64) (*domain.Inventories, error) {
	rows, err := w.db.Query(ctx, getByID, inventoryID)
	if err != nil {
		return nil, err
	}
	inventory, err := db.CollectOneRow[domain.Inventories](rows)
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

// FindDevice implements domain.IInventoryRepository.
func (w *Inventory) FindDevice(ctx context.Context, deviceSpecs domain.InventoryDeviveSpecs) (*domain.Inventories, error) {
	rows, err := w.db.Query(ctx, getAvailableProducts,
		deviceSpecs.ProductID, deviceSpecs.RAM, deviceSpecs.Ssd, deviceSpecs.Color)
	if err != nil {
		return nil, err
	}
	inventory, err := db.CollectOneRow[domain.Inventories](rows)
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
		product.ProductID, product.Price,
		string(specsjson), product.Available, product.CurrencyCode,
		"active",
	)
}
