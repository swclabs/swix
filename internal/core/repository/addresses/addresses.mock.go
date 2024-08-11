package addresses

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"

	"github.com/stretchr/testify/mock"
)

// NewAddressesMock creates a new mock object for IAddressRepository.
func NewAddressesMock() *Mock {
	return &Mock{}
}

var _ IAddressRepository = (*Mock)(nil)

// Mock is a mock type for IAddressRepository.
type Mock struct {
	mock.Mock
}

// GetByUserID implements IAddressRepository.
func (a *Mock) GetByUserID(ctx context.Context, userID int64) ([]entity.Addresses, error) {
	panic("unimplemented")
}

// Insert implements entity.IAddressRepository.
func (a *Mock) Insert(ctx context.Context, data entity.Addresses) error {
	args := a.Called(ctx, data)
	return args.Error(0)
}
