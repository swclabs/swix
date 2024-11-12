package users

import (
	"context"

	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/internal/core/domain/model"

	"github.com/stretchr/testify/mock"
)

// Mock represents a mock for IUserRepository.
type Mock struct {
	mock.Mock
}

var _ IUsers = (*Mock)(nil)

// NewUsersMock creates a new mock for IUserRepository.
func NewUsersMock() *Mock {
	return &Mock{}
}

// GetByEmail implements IUserRepository.
func (u *Mock) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	args := u.Called(ctx, email)
	return args.Get(0).(*entity.User), args.Error(1)
}

// GetByPhone implements IUserRepository.
func (u *Mock) GetByPhone(ctx context.Context, nPhone string) (*entity.User, error) {
	args := u.Called(ctx, nPhone)
	return args.Get(0).(*entity.User), args.Error(1)
}

// Info implements IUserRepository.
func (u *Mock) Info(ctx context.Context, email string) (*model.Users, error) {
	args := u.Called(ctx, email)
	return args.Get(0).(*model.Users), args.Error(1)
}

// Insert implements IUserRepository.
func (u *Mock) Insert(ctx context.Context, usr entity.User) (int64, error) {
	args := u.Called(ctx, usr)
	return args.Get(1).(int64), args.Error(0)
}

// OAuth2SaveInfo implements IUserRepository.
func (u *Mock) OAuth2SaveInfo(ctx context.Context, user entity.User) error {
	args := u.Called(ctx, user)
	return args.Error(0)
}

// Save implements IUserRepository.
func (u *Mock) Save(ctx context.Context, user entity.User) error {
	args := u.Called(ctx, user)
	return args.Error(0)
}

// GetByID implements IUserRepository.
func (u *Mock) GetByID(_ context.Context, _ int64) (*entity.User, error) {
	panic("unimplemented")
}
