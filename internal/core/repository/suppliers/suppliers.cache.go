// Package suppliers suppliers repository implementation
package suppliers

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"
)

type cache struct {
	supplier ISuppliersRepository
}

var _ ISuppliersRepository = (*cache)(nil)

func useCache(repo ISuppliersRepository) ISuppliersRepository {
	return &cache{supplier: repo}
}

// GetByPhone implements ISuppliersRepository.
func (c *cache) GetByPhone(ctx context.Context, phone string) (*entity.Suppliers, error) {
	return c.supplier.GetByPhone(ctx, phone)
}

// GetLimit implements ISuppliersRepository.
func (c *cache) GetLimit(ctx context.Context, limit int) ([]entity.Suppliers, error) {
	return c.supplier.GetLimit(ctx, limit)
}

// Insert implements ISuppliersRepository.
func (c *cache) Insert(ctx context.Context, sup entity.Suppliers) error {
	return c.supplier.Insert(ctx, sup)
}

// InsertAddress implements ISuppliersRepository.
func (c *cache) InsertAddress(ctx context.Context, addr entity.SuppliersAddress) error {
	return c.supplier.InsertAddress(ctx, addr)
}

func (c *cache) Edit(ctx context.Context, sup entity.Suppliers) error {
	return c.supplier.Edit(ctx, sup)
}
