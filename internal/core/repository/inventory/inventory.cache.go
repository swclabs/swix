package inventory

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

type cache struct {
	inventory IInventoryRepository
}

var _ IInventoryRepository = (*cache)(nil)

func useCache(repo IInventoryRepository) IInventoryRepository {
	return &cache{inventory: repo}
}

// GetById implements IInventoryRepository.
func (c *cache) GetById(ctx context.Context, inventoryId int64) (*domain.Inventory, error) {
	return c.inventory.GetById(ctx, inventoryId)
}

// GetProducts implements IInventoryRepository.
func (c *cache) GetProducts(ctx context.Context, productID string, ram string, ssd string, color string) (*domain.Inventory, error) {
	return c.inventory.GetProducts(ctx, productID, ram, ssd, color)
}

// InsertProduct implements IInventoryRepository.
func (c *cache) InsertProduct(ctx context.Context, product domain.InventoryStruct) error {
	return c.inventory.InsertProduct(ctx, product)
}
