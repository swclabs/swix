// Package inventories implement inventories
package inventories

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

var _ IInventoryRepository = (*cache)(nil)

func useCache(repo IInventoryRepository) IInventoryRepository {
	return &cache{inventory: repo}
}

type cache struct {
	inventory IInventoryRepository
}

// DeleteByID implements IInventoryRepository.
func (c *cache) DeleteByID(ctx context.Context, inventoryID int64) error {
	return c.inventory.DeleteByID(ctx, inventoryID)
}

// GetLimit implements IInventoryRepository.
func (c *cache) GetLimit(ctx context.Context, limit int, offset int) ([]domain.Inventories, error) {
	return c.inventory.GetLimit(ctx, limit, offset)
}

// GetByProductId implements IInventoryRepository.
func (c *cache) GetByProductID(ctx context.Context, ID int64) ([]domain.Inventories, error) {
	return c.inventory.GetByProductID(ctx, ID)
}

// GetById implements IInventoryRepository.
func (c *cache) GetByID(ctx context.Context, inventoryID int64) (*domain.Inventories, error) {
	return c.inventory.GetByID(ctx, inventoryID)
}

// FindDevice implements IInventoryRepository.
func (c *cache) FindDevice(ctx context.Context, deviceSpecs domain.InventoryDeviveSpecs) (*domain.Inventories, error) {
	return c.inventory.FindDevice(ctx, deviceSpecs)
}

// InsertProduct implements IInventoryRepository.
func (c *cache) InsertProduct(ctx context.Context, product domain.InventoryStruct) error {
	return c.inventory.InsertProduct(ctx, product)
}
