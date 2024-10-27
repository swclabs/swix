package inventories

import (
	"context"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/internal/core/domain/model"
	"swclabs/swipex/pkg/infra/cache"
	"swclabs/swipex/pkg/infra/db"
	"swclabs/swipex/pkg/lib/errors"
)

var _ IInventories = (*Inventory)(nil)
var _ = app.Repos(Init)

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

// UploadColorImage implements IInventories.
func (w *Inventory) UploadColorImage(ctx context.Context, ID int, url string) error {
	return w.db.SafeWrite(ctx, uploadInvItemColor, url, ID)
}

// GetByColor implements IInventories.
func (w *Inventory) GetByColor(ctx context.Context, productID int64, color string) ([]entity.Inventory, error) {
	rows, err := w.db.Query(ctx, getByColor, productID, color)
	if err != nil {
		return nil, errors.Repository("500", err)
	}
	inventories, err := db.CollectRows[entity.Inventory](rows)
	if err != nil {
		return nil, errors.Repository("500", err)
	}
	return inventories, nil
}

// GetColor implements IInventories.
func (w *Inventory) GetColor(ctx context.Context, productID int64) ([]model.ColorItem, error) {
	rows, err := w.db.Query(ctx, groupByColor, productID)
	if err != nil {
		return nil, errors.Repository("500", err)
	}
	colors, err := db.CollectRows[model.ColorItem](rows)
	if err != nil {
		return nil, errors.Repository("500", err)
	}
	return colors, nil
}

// Update implements IInventoryRepository.
func (w *Inventory) Update(ctx context.Context, inventory entity.Inventory) error {
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
func (w *Inventory) GetLimit(ctx context.Context, limit int, offset int) ([]entity.Inventory, error) {
	if offset < 1 {
		offset = 1
	}
	rows, err := w.db.Query(ctx, getProductsLimit, limit, (offset-1)*limit)
	if err != nil {
		return nil, err
	}
	return db.CollectRows[entity.Inventory](rows)
}

// GetByProductID implements IInventoryRepository.
func (w *Inventory) GetByProductID(ctx context.Context, productID int64) ([]entity.Inventory, error) {
	rows, err := w.db.Query(ctx, getByProductID, productID)
	if err != nil {
		return nil, errors.Repository("500", err)
	}
	inventories, err := db.CollectRows[entity.Inventory](rows)
	if err != nil {
		return nil, errors.Repository("500", err)
	}
	return inventories, nil
}

// GetByID implements IInventoryRepository.
func (w *Inventory) GetByID(ctx context.Context, inventoryID int64) (*entity.Inventory, error) {
	rows, err := w.db.Query(ctx, getByID, inventoryID)
	if err != nil {
		return nil, err
	}
	inventory, err := db.CollectRow[entity.Inventory](rows)
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

// InsertProduct implements IInventoryRepository.
func (w *Inventory) InsertProduct(ctx context.Context, product entity.Inventory) (int64, error) {
	return w.db.SafeWriteReturn(ctx, insertIntoInventory,
		product.ProductID, product.Price,
		product.Available, product.CurrencyCode,
		product.Status, product.Image, product.Color, product.ColorImg, product.Specs,
	)
}
