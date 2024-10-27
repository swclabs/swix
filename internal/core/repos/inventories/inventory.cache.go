// Package inventories implement inventories
package inventories

import (
	"context"
	"fmt"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/internal/core/domain/model"
	"swclabs/swipex/pkg/infra/cache"
	"swclabs/swipex/pkg/lib/crypto"
)

var (
	_                 IInventories = (*_Cache)(nil)
	keyGetByID                     = "IInventoryRepository:GetByID:%d"
	keyGetByProductID              = "IInventoryRepository:GetByProductID:%d"
)

func useCache(cache cache.ICache, repo IInventories) IInventories {
	return &_Cache{
		inventory: repo,
		cache:     cache,
	}
}

type _Cache struct {
	cache     cache.ICache
	inventory IInventories
}

// UploadColorImage implements IInventories.
func (c *_Cache) UploadColorImage(ctx context.Context, ID int, url string) error {
	if err := c.inventory.UploadColorImage(ctx, ID, url); err != nil {
		return err
	}
	key := crypto.HashOf(fmt.Sprintf(keyGetByID, ID))
	return cache.Delete(ctx, c.cache, key)
}

// GetByColor implements IInventories.
func (c *_Cache) GetByColor(ctx context.Context, productID int64, color string) ([]entity.Inventory, error) {
	return c.inventory.GetByColor(ctx, productID, color)
}

// GetColor implements IInventories.
func (c *_Cache) GetColor(ctx context.Context, productID int64) ([]model.ColorItem, error) {
	return c.inventory.GetColor(ctx, productID)
}

// Update implements IInventoryRepository.
func (c *_Cache) Update(ctx context.Context, inventory entity.Inventory) error {
	if err := c.inventory.Update(ctx, inventory); err != nil {
		return err
	}
	key := crypto.HashOf(fmt.Sprintf(keyGetByID, inventory.ID))
	return cache.Delete(ctx, c.cache, key)
}

// UploadImage implements IInventoryRepository.
func (c *_Cache) UploadImage(ctx context.Context, ID int, url string) error {
	if err := c.inventory.UploadImage(ctx, ID, url); err != nil {
		return err
	}
	key := crypto.HashOf(fmt.Sprintf(keyGetByID, ID))
	return cache.Delete(ctx, c.cache, key)
}

// DeleteByID implements IInventoryRepository.
func (c *_Cache) DeleteByID(ctx context.Context, inventoryID int64) error {
	if err := c.inventory.DeleteByID(ctx, inventoryID); err != nil {
		return err
	}
	key := crypto.HashOf(fmt.Sprintf(keyGetByID, inventoryID))
	return cache.Delete(ctx, c.cache, key)
}

// InsertProduct implements IInventoryRepository.
func (c *_Cache) InsertProduct(ctx context.Context, product entity.Inventory) (int64, error) {
	ID, err := c.inventory.InsertProduct(ctx, product)
	if err != nil {
		return -1, err
	}
	key := crypto.HashOf(fmt.Sprintf(keyGetByProductID, product.ProductID))
	return ID, cache.Delete(ctx, c.cache, key)
}

// GetLimit implements IInventoryRepository.
func (c *_Cache) GetLimit(ctx context.Context, limit int, offset int) ([]entity.Inventory, error) {
	key := crypto.HashOf(fmt.Sprintf("IInventoryRepository:GetLimit:%d:%d", limit, offset))
	result, err := cache.GetSlice[entity.Inventory](ctx, c.cache, key)
	if err != nil {
		result, err = c.inventory.GetLimit(ctx, limit, offset)
		if err != nil {
			return nil, err
		}
		if err := cache.Set[[]entity.Inventory](ctx, c.cache, key, result); err != nil {
			return result, nil
		}
	}
	return result, nil
}

// GetByProductID implements IInventoryRepository.
func (c *_Cache) GetByProductID(ctx context.Context, ID int64) ([]entity.Inventory, error) {
	key := crypto.HashOf(fmt.Sprintf(keyGetByProductID, ID))
	result, err := cache.GetSlice[entity.Inventory](ctx, c.cache, key)
	if err != nil {
		result, err = c.inventory.GetByProductID(ctx, ID)
		if err != nil {
			return nil, err
		}
		if err := cache.Set[[]entity.Inventory](ctx, c.cache, key, result); err != nil {
			return result, nil
		}
	}
	return result, nil
}

// GetByID implements IInventoryRepository.
func (c *_Cache) GetByID(ctx context.Context, inventoryID int64) (*entity.Inventory, error) {
	key := crypto.HashOf(fmt.Sprintf(keyGetByID, inventoryID))
	result, err := cache.Get[entity.Inventory](ctx, c.cache, key)
	if err != nil {
		result, err = c.inventory.GetByID(ctx, inventoryID)
		if err != nil {
			return nil, err
		}
		if err := cache.Set[entity.Inventory](ctx, c.cache, key, *result); err != nil {
			return result, nil
		}
	}
	return result, nil
}
