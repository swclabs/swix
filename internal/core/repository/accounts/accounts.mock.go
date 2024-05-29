package accounts

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type Mock struct {
	mock.Mock
}

var _ IAccountRepository = (*Mock)(nil)

func NewAccountsMock() *Mock {
	return &Mock{}
}

// GetByEmail implements domain.IAccountRepository.
func (a *Mock) GetByEmail(ctx context.Context, email string) (*domain.Account, error) {
	panic("unimplemented")
}

// Insert implements domain.IAccountRepository.
func (a *Mock) Insert(ctx context.Context, acc *domain.Account) error {
	panic("unimplemented")
}

// SaveInfo implements domain.IAccountRepository.
func (a *Mock) SaveInfo(ctx context.Context, acc *domain.Account) error {
	panic("unimplemented")
}

// Use implements domain.IAccountRepository.
func (a *Mock) Use(tx *gorm.DB) IAccountRepository {
	panic("unimplemented")
}
