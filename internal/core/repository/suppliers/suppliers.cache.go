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
	panic("unimplemented")
}

// GetLimit implements ISuppliersRepository.
func (c *cache) GetLimit(ctx context.Context, limit int) ([]domain.Suppliers, error) {
	panic("unimplemented")
}

// Insert implements ISuppliersRepository.
func (c *cache) Insert(ctx context.Context, sup domain.Suppliers, addr domain.Addresses) error {
	panic("unimplemented")
}

// InsertAddress implements ISuppliersRepository.
func (c *cache) InsertAddress(ctx context.Context, addr domain.SuppliersAddress) error {
	panic("unimplemented")
}
