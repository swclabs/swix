package inventories

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

// GetLimit implements IInventoryRepository.
func (c *cache) GetLimit(ctx context.Context, limit int, offset int) ([]domain.Inventories, error) {
	return c.inventory.GetLimit(ctx, limit, offset)
}

// GetByProductId implements IInventoryRepository.
func (c *cache) GetByProductId(ctx context.Context, productId int64) ([]domain.Inventories, error) {
	//TODO implement me
	panic("implement me")
}

// GetById implements IInventoryRepository.
func (c *cache) GetById(ctx context.Context, inventoryId int64) (*domain.Inventories, error) {
	return c.inventory.GetById(ctx, inventoryId)
}

// FindDevice implements IInventoryRepository.
func (c *cache) FindDevice(ctx context.Context, productID string, ram string, ssd string, color string) (*domain.Inventories, error) {
	return c.inventory.FindDevice(ctx, productID, ram, ssd, color)
}

// InsertProduct implements IInventoryRepository.
func (c *cache) InsertProduct(ctx context.Context, product domain.InventoryStruct) error {
	return c.inventory.InsertProduct(ctx, product)
}
