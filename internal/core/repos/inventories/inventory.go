package inventories

import (
	"context"
	"swclabs/swix/boot"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/pkg/infra/cache"
	"swclabs/swix/pkg/infra/db"
	"swclabs/swix/pkg/lib/errors"
)

var _ IInventories = (*Inventory)(nil)
var _ = boot.Repos(Init)

// New creates a new Inventory object.
func New(conn db.IDatabase) IInventories {
	return &Inventory{
		db: conn,
	}
}

// Init initializes the Inventory object with database and redis connection
func Init(conn db.IDatabase, cache cache.ICache) IInventories {
	return useCache(cache, &Inventory{db: conn})
}

// Inventory represents the Inventory object.
type Inventory struct {
	db db.IDatabase
}

// Update implements IInventoryRepository.
func (w *Inventory) Update(ctx context.Context, inventory entity.Inventories) error {
	return w.db.SafeWrite(ctx, update,
		inventory.ID,
		inventory.ProductID,
		inventory.Status,
		inventory.Price.String(),
		inventory.CurrencyCode,
		inventory.Available,
		inventory.Image,
		inventory.Color,
		inventory.ColorImg,
	)
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
func (w *Inventory) GetLimit(ctx context.Context, limit int, offset int) ([]entity.Inventories, error) {
	if offset < 1 {
		offset = 1
	}
	rows, err := w.db.Query(ctx, getProductsLimit, limit, (offset-1)*limit)
	if err != nil {
		return nil, err
	}
	return db.CollectRows[entity.Inventories](rows)
}

// GetByProductID implements IInventoryRepository.
func (w *Inventory) GetByProductID(ctx context.Context, productID int64) ([]entity.Inventories, error) {
	rows, err := w.db.Query(ctx, getByProductID, productID)
	if err != nil {
		return nil, errors.Repository("500", err)
	}
	inventories, err := db.CollectRows[entity.Inventories](rows)
	if err != nil {
		return nil, errors.Repository("500", err)
	}
	return inventories, nil
}

// GetByID implements IInventoryRepository.
func (w *Inventory) GetByID(ctx context.Context, inventoryID int64) (*entity.Inventories, error) {
	rows, err := w.db.Query(ctx, getByID, inventoryID)
	if err != nil {
		return nil, err
	}
	inventory, err := db.CollectOneRow[entity.Inventories](rows)
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

// InsertProduct implements IInventoryRepository.
func (w *Inventory) InsertProduct(ctx context.Context, product entity.Inventories) (int64, error) {
	return w.db.SafeWriteReturn(ctx, insertIntoInventory,
		product.ProductID, product.Price,
		product.Available, product.CurrencyCode,
		product.Status, product.Image, product.Color, product.ColorImg,
	)
}
