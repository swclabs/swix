package repository

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type AccountMock struct {
	mock.Mock
}

var _ domain.IAccountRepository = (*AccountMock)(nil)

func NewAccountsMock() *AccountMock {
	return &AccountMock{}
}

// GetByEmail implements domain.IAccountRepository.
func (a *AccountMock) GetByEmail(ctx context.Context, email string) (*domain.Account, error) {
	panic("unimplemented")
}

// Insert implements domain.IAccountRepository.
func (a *AccountMock) Insert(ctx context.Context, acc *domain.Account) error {
	panic("unimplemented")
}

// SaveInfo implements domain.IAccountRepository.
func (a *AccountMock) SaveInfo(ctx context.Context, acc *domain.Account) error {
	panic("unimplemented")
}

// Use implements domain.IAccountRepository.
func (a *AccountMock) Use(tx *gorm.DB) domain.IAccountRepository {
	panic("unimplemented")
}
