package suppliers

import (
	"context"
	"swclabs/swipex/internal/core/domain/entity"

	"github.com/stretchr/testify/mock"
)

// Mock is a mock for ISuppliersRepository.
type Mock struct {
	mock.Mock
}

var _ ISuppliers = (*Mock)(nil)

// GetByPhone implements ISuppliersRepository.
func (s *Mock) GetByPhone(ctx context.Context, email string) (*entity.Supplier, error) {
	args := s.Called(ctx, email)
	return args.Get(0).(*entity.Supplier), args.Error(1)
}

// GetLimit implements ISuppliersRepository.
func (s *Mock) GetLimit(ctx context.Context, limit int) ([]entity.Supplier, error) {
	args := s.Called(ctx, limit)
	return args.Get(0).([]entity.Supplier), args.Error(1)
}

// Insert implements ISuppliersRepository.
func (s *Mock) Insert(ctx context.Context, sup entity.Supplier) error {
	args := s.Called(ctx, sup)
	return args.Error(0)
}

// Edit implements ISuppliersRepository.
func (s *Mock) Edit(ctx context.Context, sup entity.Supplier) error {
	args := s.Called(ctx, sup)
	return args.Error(0)
}
