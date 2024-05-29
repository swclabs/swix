package users

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type UsersMock struct {
	mock.Mock
}

var _ IUserRepository = (*UsersMock)(nil)

func NewUsersMock() *UsersMock {
	return &UsersMock{}
}

// GetByEmail implements domain.IUserRepository.
func (u *UsersMock) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	panic("unimplemented")
}

// GetByPhone implements domain.IUserRepository.
func (u *UsersMock) GetByPhone(ctx context.Context, nPhone string) (*domain.User, error) {
	panic("unimplemented")
}

// Info implements domain.IUserRepository.
func (u *UsersMock) Info(ctx context.Context, email string) (*domain.UserInfo, error) {
	panic("unimplemented")
}

// Insert implements domain.IUserRepository.
func (u *UsersMock) Insert(ctx context.Context, usr *domain.User) error {
	panic("unimplemented")
}

// OAuth2SaveInfo implements domain.IUserRepository.
func (u *UsersMock) OAuth2SaveInfo(ctx context.Context, user *domain.User) error {
	panic("unimplemented")
}

// SaveInfo implements domain.IUserRepository.
func (u *UsersMock) SaveInfo(ctx context.Context, user *domain.User) error {
	panic("unimplemented")
}

// TransactionSaveOAuth2 implements domain.IUserRepository.
func (u *UsersMock) TransactionSaveOAuth2(ctx context.Context, data *domain.User) error {
	panic("unimplemented")
}

// TransactionSignUp implements domain.IUserRepository.
func (u *UsersMock) TransactionSignUp(ctx context.Context, user *domain.User, password string) error {
	panic("unimplemented")
}

// UpdateProperties implements domain.IUserRepository.
func (u *UsersMock) UpdateProperties(ctx context.Context, query string, user *domain.User) error {
	panic("unimplemented")
}

// Use implements domain.IUserRepository.
func (u *UsersMock) Use(tx *gorm.DB) IUserRepository {
	panic("unimplemented")
}
