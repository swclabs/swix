package addresses

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type AddressesMock struct {
	mock.Mock
}

var _ IAddressRepository = (*AddressesMock)(nil)

func NewAddressesMock() *AddressesMock {
	return &AddressesMock{}
}

// Insert implements domain.IAddressRepository.
func (a *AddressesMock) Insert(ctx context.Context, data *domain.Addresses) error {
	panic("unimplemented")
}

// Use implements domain.IAddressRepository.
func (a *AddressesMock) Use(tx *gorm.DB) IAddressRepository {
	panic("unimplemented")
}
