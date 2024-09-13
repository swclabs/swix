package accounts

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"

	"github.com/stretchr/testify/mock"
)

// Mock is a mock for IAccountRepository.
type Mock struct {
	mock.Mock
}

var _ IAccounts = (*Mock)(nil)

// NewAccountsMock returns a new Mock object
func NewAccountsMock() *Mock {
	return &Mock{}
}

// GetByEmail implements IAccountRepository.
func (a *Mock) GetByEmail(ctx context.Context, email string) (*entity.Account, error) {
	args := a.Called(ctx, email)
	return args.Get(0).(*entity.Account), args.Error(0)
}

// Insert implements IAccountRepository.
func (a *Mock) Insert(ctx context.Context, acc entity.Account) error {
	args := a.Called(ctx, acc)
	return args.Error(0)
}

// SaveInfo implements IAccountRepository.
func (a *Mock) SaveInfo(ctx context.Context, acc entity.Account) error {
	args := a.Called(ctx, acc)
	return args.Error(0)
}
