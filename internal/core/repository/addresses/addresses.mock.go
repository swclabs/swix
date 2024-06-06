package addresses

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

var _ IAddressRepository = (*Mock)(nil)

func NewAddressesMock() *Mock {
	return &Mock{}
}

// Insert implements domain.IAddressRepository.
func (a *Mock) Insert(ctx context.Context, data *domain.Addresses) error {
	args := a.Called(ctx, data)
	return args.Error(0)
}
