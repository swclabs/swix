package addresses

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"

	"github.com/stretchr/testify/mock"
)

// Mock is a mock type for IAddressRepository.
type Mock struct {
	mock.Mock
}

var _ IAddressRepository = (*Mock)(nil)

// NewAddressesMock creates a new mock object for IAddressRepository.
func NewAddressesMock() *Mock {
	return &Mock{}
}

// Insert implements entity.IAddressRepository.
func (a *Mock) Insert(ctx context.Context, data entity.Addresses) error {
	args := a.Called(ctx, data)
	return args.Error(0)
}
