// Package suppliers suppliers repos implementation
package suppliers

import (
	"context"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/pkg/infra/cache"
)

type _cache struct {
	cache    cache.ICache
	supplier ISuppliers
}

var _ ISuppliers = (*_cache)(nil)

func useCache(cache cache.ICache, repo ISuppliers) ISuppliers {
	return &_cache{
		supplier: repo,
		cache:    cache,
	}
}

// GetByPhone implements ISuppliersRepository.
func (c *_cache) GetByPhone(ctx context.Context, phone string) (*entity.Supplier, error) {
	return c.supplier.GetByPhone(ctx, phone)
}

// GetLimit implements ISuppliersRepository.
func (c *_cache) GetLimit(ctx context.Context, limit int) ([]entity.Supplier, error) {
	return c.supplier.GetLimit(ctx, limit)
}

// Insert implements ISuppliersRepository.
func (c *_cache) Insert(ctx context.Context, sup entity.Supplier) error {
	return c.supplier.Insert(ctx, sup)
}

// Edit implements ISuppliersRepository.
func (c *_cache) Edit(ctx context.Context, sup entity.Supplier) error {
	return c.supplier.Edit(ctx, sup)
}
