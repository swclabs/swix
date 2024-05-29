package addresses

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
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
	panic("unimplemented")
}

// Use implements domain.IAddressRepository.
func (a *Mock) Use(tx *gorm.DB) IAddressRepository {
	panic("unimplemented")
}
