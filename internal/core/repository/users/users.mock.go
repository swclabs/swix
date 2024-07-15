package users

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

// Mock represents a mock for IUserRepository.
type Mock struct {
	mock.Mock
}

var _ IUserRepository = (*Mock)(nil)

// NewUsersMock creates a new mock for IUserRepository.
func NewUsersMock() *Mock {
	return &Mock{}
}

// GetByEmail implements domain.IUserRepository.
func (u *Mock) GetByEmail(ctx context.Context, email string) (*domain.Users, error) {
	args := u.Called(ctx, email)
	return args.Get(0).(*domain.Users), args.Error(1)
}

// GetByPhone implements domain.IUserRepository.
func (u *Mock) GetByPhone(ctx context.Context, nPhone string) (*domain.Users, error) {
	args := u.Called(ctx, nPhone)
	return args.Get(0).(*domain.Users), args.Error(1)
}

// Info implements domain.IUserRepository.
func (u *Mock) Info(ctx context.Context, email string) (*domain.UserSchema, error) {
	args := u.Called(ctx, email)
	return args.Get(0).(*domain.UserSchema), args.Error(1)
}

// Insert implements domain.IUserRepository.
func (u *Mock) Insert(ctx context.Context, usr domain.Users) error {
	args := u.Called(ctx, usr)
	return args.Error(0)
}

// OAuth2SaveInfo implements domain.IUserRepository.
func (u *Mock) OAuth2SaveInfo(ctx context.Context, user domain.Users) error {
	args := u.Called(ctx, user)
	return args.Error(0)
}

// SaveInfo implements domain.IUserRepository.
func (u *Mock) SaveInfo(ctx context.Context, user domain.Users) error {
	args := u.Called(ctx, user)
	return args.Error(0)
}

// UpdateProperties implements domain.IUserRepository.
func (u *Mock) UpdateProperties(ctx context.Context, query string, user domain.Users) error {
	args := u.Called(ctx, query, user)
	return args.Error(0)
}

// GetByID implements IUserRepository.
func (u *Mock) GetByID(ctx context.Context, id int64) (*domain.Users, error) {
	panic("unimplemented")
}
