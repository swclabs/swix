package suppliers

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

// Mock is a mock for ISuppliersRepository.
type Mock struct {
	mock.Mock
}

var _ ISuppliersRepository = (*Mock)(nil)

// GetByPhone implements domain.ISuppliersRepository.
func (s *Mock) GetByPhone(ctx context.Context, email string) (*domain.Suppliers, error) {
	args := s.Called(ctx, email)
	return args.Get(0).(*domain.Suppliers), args.Error(1)
}

// GetLimit implements domain.ISuppliersRepository.
func (s *Mock) GetLimit(ctx context.Context, limit int) ([]domain.Suppliers, error) {
	args := s.Called(ctx, limit)
	return args.Get(0).([]domain.Suppliers), args.Error(1)
}

// Insert implements domain.ISuppliersRepository.
func (s *Mock) Insert(ctx context.Context, sup domain.Suppliers) error {
	args := s.Called(ctx, sup)
	return args.Error(0)
}

// InsertAddress implements domain.ISuppliersRepository.
func (s *Mock) InsertAddress(ctx context.Context, addr domain.SuppliersAddress) error {
	args := s.Called(ctx, addr)
	return args.Error(0)
}

// Edit implements domain.ISuppliersRepository.
func (s *Mock) Edit(ctx context.Context, sup domain.Suppliers) error {
	args := s.Called(ctx, sup)
	return args.Error(0)
}
