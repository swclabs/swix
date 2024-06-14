package accounts

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
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
	args := a.Called(ctx, email)
	return args.Get(0).(*domain.Account), args.Error(0)
}

// Insert implements domain.IAccountRepository.
func (a *Mock) Insert(ctx context.Context, acc domain.Account) error {
	args := a.Called(ctx, acc)
	return args.Error(0)
}

// SaveInfo implements domain.IAccountRepository.
func (a *Mock) SaveInfo(ctx context.Context, acc domain.Account) error {
	args := a.Called(ctx, acc)
	return args.Error(0)
}
