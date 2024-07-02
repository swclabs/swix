package suppliers

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

type cache struct {
	supplier ISuppliersRepository
}

var _ ISuppliersRepository = (*cache)(nil)

func useCache(repo ISuppliersRepository) ISuppliersRepository {
	return &cache{supplier: repo}
}

// GetByPhone implements ISuppliersRepository.
func (c *cache) GetByPhone(ctx context.Context, phone string) (*domain.Suppliers, error) {
	return c.supplier.GetByPhone(ctx, phone)
}

// GetLimit implements ISuppliersRepository.
func (c *cache) GetLimit(ctx context.Context, limit int) ([]domain.Suppliers, error) {
	return c.supplier.GetLimit(ctx, limit)
}

// Insert implements ISuppliersRepository.
func (c *cache) Insert(ctx context.Context, sup domain.Suppliers, addr domain.Addresses) error {
	return c.supplier.Insert(ctx, sup, addr)
}

// InsertAddress implements ISuppliersRepository.
func (c *cache) InsertAddress(ctx context.Context, addr domain.SuppliersAddress) error {
	return c.supplier.InsertAddress(ctx, addr)
}
