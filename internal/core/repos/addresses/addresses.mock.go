package addresses

import (
	"context"
	"swclabs/swipex/internal/core/domain/entity"

	"github.com/stretchr/testify/mock"
)

// NewAddressesMock creates a new mock object for IAddressRepository.
func NewAddressesMock() *Mock {
	return &Mock{}
}

var _ IAddress = (*Mock)(nil)

// Mock is a mock type for IAddressRepository.
type Mock struct {
	mock.Mock
}

// GetByID implements IAddress.
func (a *Mock) GetByID(ctx context.Context, id int64) (*entity.Address, error) {
	panic("unimplemented")
}

// GetByUserID implements IAddressRepository.
func (a *Mock) GetByUserID(_ context.Context, _ int64) ([]entity.Address, error) {
	panic("unimplemented")
}

// Insert implements entity.IAddressRepository.
func (a *Mock) Insert(ctx context.Context, data entity.Address) (int64, error) {
	args := a.Called(ctx, data)
	return args.Get(0).(int64), args.Error(1)
}
