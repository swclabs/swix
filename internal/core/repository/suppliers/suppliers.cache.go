// Package suppliers suppliers repository implementation
package suppliers

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/pkg/infra/cache"
)

type _cache struct {
	cache    cache.ICache
	supplier ISuppliersRepository
}

var _ ISuppliersRepository = (*_cache)(nil)

func useCache(cache cache.ICache, repo ISuppliersRepository) ISuppliersRepository {
	return &_cache{
		supplier: repo,
		cache:    cache,
	}
}

// GetByPhone implements ISuppliersRepository.
func (c *_cache) GetByPhone(ctx context.Context, phone string) (*entity.Suppliers, error) {
	return c.supplier.GetByPhone(ctx, phone)
}

// GetLimit implements ISuppliersRepository.
func (c *_cache) GetLimit(ctx context.Context, limit int) ([]entity.Suppliers, error) {
	return c.supplier.GetLimit(ctx, limit)
}

// Insert implements ISuppliersRepository.
func (c *_cache) Insert(ctx context.Context, sup entity.Suppliers) error {
	return c.supplier.Insert(ctx, sup)
}

// InsertAddress implements ISuppliersRepository.
func (c *_cache) InsertAddress(ctx context.Context, addr entity.SuppliersAddress) error {
	return c.supplier.InsertAddress(ctx, addr)
}

func (c *_cache) Edit(ctx context.Context, sup entity.Suppliers) error {
	return c.supplier.Edit(ctx, sup)
}
