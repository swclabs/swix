package users

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

var _ IUserRepository = (*Mock)(nil)

func NewUsersMock() *Mock {
	return &Mock{}
}

// GetByEmail implements domain.IUserRepository.
func (u *Mock) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	args := u.Called(ctx, email)
	return args.Get(0).(*domain.User), args.Error(1)
}

// GetByPhone implements domain.IUserRepository.
func (u *Mock) GetByPhone(ctx context.Context, nPhone string) (*domain.User, error) {
	args := u.Called(ctx, nPhone)
	return args.Get(0).(*domain.User), args.Error(1)
}

// Info implements domain.IUserRepository.
func (u *Mock) Info(ctx context.Context, email string) (*domain.UserInfo, error) {
	args := u.Called(ctx, email)
	return args.Get(0).(*domain.UserInfo), args.Error(1)
}

// Insert implements domain.IUserRepository.
func (u *Mock) Insert(ctx context.Context, usr domain.User) error {
	args := u.Called(ctx, usr)
	return args.Error(0)
}

// OAuth2SaveInfo implements domain.IUserRepository.
func (u *Mock) OAuth2SaveInfo(ctx context.Context, user domain.User) error {
	args := u.Called(ctx, user)
	return args.Error(0)
}

// SaveInfo implements domain.IUserRepository.
func (u *Mock) SaveInfo(ctx context.Context, user domain.User) error {
	args := u.Called(ctx, user)
	return args.Error(0)
}

// TransactionSaveOAuth2 implements domain.IUserRepository.
func (u *Mock) TransactionSaveOAuth2(ctx context.Context, data domain.User) error {
	args := u.Called(ctx, data)
	return args.Error(0)
}

// TransactionSignUp implements domain.IUserRepository.
func (u *Mock) TransactionSignUp(ctx context.Context, user domain.User, password string) error {
	args := u.Called(ctx, user, password)
	return args.Error(0)
}

// UpdateProperties implements domain.IUserRepository.
func (u *Mock) UpdateProperties(ctx context.Context, query string, user domain.User) error {
	args := u.Called(ctx, query, user)
	return args.Error(0)
}
