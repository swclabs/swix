// Package inventories implement inventories
package inventories

import (
	"context"
	"fmt"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/pkg/infra/cache"
	"swclabs/swix/pkg/lib/crypto"
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

// Update implements IInventoryRepository.
func (c *_Cache) Update(ctx context.Context, inventory entity.Inventories) error {
	if err := c.inventory.Update(ctx, inventory); err != nil {
		return err
	}
	key := crypto.HashOf(fmt.Sprintf(keyGetByID, inventory.ID))
	return cache.Set(ctx, c.cache, key, inventory)
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
func (c *_Cache) InsertProduct(ctx context.Context, product entity.Inventories) (int64, error) {
	ID, err := c.inventory.InsertProduct(ctx, product)
	if err != nil {
		return -1, err
	}
	key := crypto.HashOf(fmt.Sprintf(keyGetByProductID, product.ProductID))
	return ID, cache.Delete(ctx, c.cache, key)
}

// GetLimit implements IInventoryRepository.
func (c *_Cache) GetLimit(ctx context.Context, limit int, offset int) ([]entity.Inventories, error) {
	key := crypto.HashOf(fmt.Sprintf("IInventoryRepository:GetLimit:%d:%d", limit, offset))
	result, err := cache.GetSlice[entity.Inventories](ctx, c.cache, key)
	if err != nil {
		result, err = c.inventory.GetLimit(ctx, limit, offset)
		if err != nil {
			return nil, err
		}
		if err := cache.Set[[]entity.Inventories](ctx, c.cache, key, result); err != nil {
			return result, nil
		}
	}
	return result, nil
}

// GetByProductID implements IInventoryRepository.
func (c *_Cache) GetByProductID(ctx context.Context, ID int64) ([]entity.Inventories, error) {
	key := crypto.HashOf(fmt.Sprintf(keyGetByProductID, ID))
	result, err := cache.GetSlice[entity.Inventories](ctx, c.cache, key)
	if err != nil {
		result, err = c.inventory.GetByProductID(ctx, ID)
		if err != nil {
			return nil, err
		}
		if err := cache.Set[[]entity.Inventories](ctx, c.cache, key, result); err != nil {
			return result, nil
		}
	}
	return result, nil
}

// GetByID implements IInventoryRepository.
func (c *_Cache) GetByID(ctx context.Context, inventoryID int64) (*entity.Inventories, error) {
	key := crypto.HashOf(fmt.Sprintf(keyGetByID, inventoryID))
	result, err := cache.Get[entity.Inventories](ctx, c.cache, key)
	if err != nil {
		result, err = c.inventory.GetByID(ctx, inventoryID)
		if err != nil {
			return nil, err
		}
		if err := cache.Set[entity.Inventories](ctx, c.cache, key, *result); err != nil {
			return result, nil
		}
	}
	return result, nil
}
