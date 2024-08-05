package suppliers

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"

	"github.com/stretchr/testify/mock"
)

// Mock is a mock for ISuppliersRepository.
type Mock struct {
	mock.Mock
}

var _ ISuppliersRepository = (*Mock)(nil)

// GetByPhone implements ISuppliersRepository.
func (s *Mock) GetByPhone(ctx context.Context, email string) (*entity.Suppliers, error) {
	args := s.Called(ctx, email)
	return args.Get(0).(*entity.Suppliers), args.Error(1)
}

// GetLimit implements ISuppliersRepository.
func (s *Mock) GetLimit(ctx context.Context, limit int) ([]entity.Suppliers, error) {
	args := s.Called(ctx, limit)
	return args.Get(0).([]entity.Suppliers), args.Error(1)
}

// Insert implements ISuppliersRepository.
func (s *Mock) Insert(ctx context.Context, sup entity.Suppliers) error {
	args := s.Called(ctx, sup)
	return args.Error(0)
}

// InsertAddress implements ISuppliersRepository.
func (s *Mock) InsertAddress(ctx context.Context, addr entity.SuppliersAddress) error {
	args := s.Called(ctx, addr)
	return args.Error(0)
}

// Edit implements ISuppliersRepository.
func (s *Mock) Edit(ctx context.Context, sup entity.Suppliers) error {
	args := s.Called(ctx, sup)
	return args.Error(0)
}
