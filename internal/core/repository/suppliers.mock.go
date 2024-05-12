package repository

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type SuppliersMock struct {
	mock.Mock
}

var _ domain.ISuppliersRepository = (*SuppliersMock)(nil)

// GetByPhone implements domain.ISuppliersRepository.
func (s *SuppliersMock) GetByPhone(ctx context.Context, email string) (*domain.Suppliers, error) {
	panic("unimplemented")
}

// GetLimit implements domain.ISuppliersRepository.
func (s *SuppliersMock) GetLimit(ctx context.Context, limit int) ([]domain.Suppliers, error) {
	args := s.Called(ctx, limit)
	return args.Get(0).([]domain.Suppliers), args.Error(1)
}

// Insert implements domain.ISuppliersRepository.
func (s *SuppliersMock) Insert(ctx context.Context, sup domain.Suppliers, addr domain.Addresses) error {
	panic("unimplemented")
}

// InsertAddress implements domain.ISuppliersRepository.
func (s *SuppliersMock) InsertAddress(ctx context.Context, addr domain.SuppliersAddress) error {
	panic("unimplemented")
}

// Use implements domain.ISuppliersRepository.
func (s *SuppliersMock) Use(tx *gorm.DB) domain.ISuppliersRepository {
	panic("unimplemented")
}
