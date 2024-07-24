// Package inventories implement inventories
package inventories

import (
	"context"
	"swclabs/swipecore/internal/core/domain/dto"
	"swclabs/swipecore/internal/core/domain/entity"
)

var _ IInventoryRepository = (*cache)(nil)

func useCache(repo IInventoryRepository) IInventoryRepository {
	return &cache{inventory: repo}
}

type cache struct {
	inventory IInventoryRepository
}

// Update implements IInventoryRepository.
func (c *cache) Update(ctx context.Context, inventory entity.Inventories) error {
	return c.inventory.Update(ctx, inventory)
}

// UploadImage implements IInventoryRepository.
func (c *cache) UploadImage(ctx context.Context, ID int, url string) error {
	return c.inventory.UploadImage(ctx, ID, url)
}

// DeleteByID implements IInventoryRepository.
func (c *cache) DeleteByID(ctx context.Context, inventoryID int64) error {
	return c.inventory.DeleteByID(ctx, inventoryID)
}

// GetLimit implements IInventoryRepository.
func (c *cache) GetLimit(ctx context.Context, limit int, offset int) ([]entity.Inventories, error) {
	return c.inventory.GetLimit(ctx, limit, offset)
}

// GetByProductID implements IInventoryRepository.
func (c *cache) GetByProductID(ctx context.Context, ID int64) ([]entity.Inventories, error) {
	return c.inventory.GetByProductID(ctx, ID)
}

// GetByID implements IInventoryRepository.
func (c *cache) GetByID(ctx context.Context, inventoryID int64) (*entity.Inventories, error) {
	return c.inventory.GetByID(ctx, inventoryID)
}

// FindDevice implements IInventoryRepository.
func (c *cache) FindDevice(ctx context.Context, device dto.InventoryDeviceSpecs) (*entity.Inventories, error) {
	return c.inventory.FindDevice(ctx, device)
}

// InsertProduct implements IInventoryRepository.
func (c *cache) InsertProduct(ctx context.Context, product entity.Inventories) error {
	return c.inventory.InsertProduct(ctx, product)
}
