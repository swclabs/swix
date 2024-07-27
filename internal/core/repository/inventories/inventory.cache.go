// Package inventories implement inventories
package inventories

import (
	"context"
	"encoding/json"
	"fmt"
	"swclabs/swipecore/internal/core/domain/dtos"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/pkg/infra/cache"
	"swclabs/swipecore/pkg/lib/crypto"
)

var _ IInventoryRepository = (*_Cache)(nil)

func useCache(cache cache.ICache, repo IInventoryRepository) IInventoryRepository {
	return &_Cache{
		inventory: repo,
		cache:     cache,
	}
}

type _Cache struct {
	cache     cache.ICache
	inventory IInventoryRepository
}

// Update implements IInventoryRepository.
func (c *_Cache) Update(ctx context.Context, inventory entity.Inventories) error {
	return c.inventory.Update(ctx, inventory)
}

// UploadImage implements IInventoryRepository.
func (c *_Cache) UploadImage(ctx context.Context, ID int, url string) error {
	return c.inventory.UploadImage(ctx, ID, url)
}

// DeleteByID implements IInventoryRepository.
func (c *_Cache) DeleteByID(ctx context.Context, inventoryID int64) error {
	return c.inventory.DeleteByID(ctx, inventoryID)
}

// InsertProduct implements IInventoryRepository.
func (c *_Cache) InsertProduct(ctx context.Context, product entity.Inventories) error {
	return c.inventory.InsertProduct(ctx, product)
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
	key := crypto.HashOf(fmt.Sprintf("IInventoryRepository:GetByProductID:%d", ID))
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
	key := crypto.HashOf(fmt.Sprintf("IInventoryRepository:GetByID:%d", inventoryID))
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

// FindDevice implements IInventoryRepository.
func (c *_Cache) FindDevice(ctx context.Context, device dtos.InventoryDeviceSpecs) (*entity.Inventories, error) {
	data, _ := json.Marshal(device)
	key := crypto.HashOf(fmt.Sprintf("IInventoryRepository:FindDevice:%s", string(data)))
	result, err := cache.Get[entity.Inventories](ctx, c.cache, key)
	if err != nil {
		result, err = c.inventory.FindDevice(ctx, device)
		if err != nil {
			return nil, err
		}
		if err := cache.Set[entity.Inventories](ctx, c.cache, key, *result); err != nil {
			return result, nil
		}
	}
	return result, nil
}
