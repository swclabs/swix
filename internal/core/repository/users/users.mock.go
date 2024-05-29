package users

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
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
	panic("unimplemented")
}

// GetByPhone implements domain.IUserRepository.
func (u *Mock) GetByPhone(ctx context.Context, nPhone string) (*domain.User, error) {
	panic("unimplemented")
}

// Info implements domain.IUserRepository.
func (u *Mock) Info(ctx context.Context, email string) (*domain.UserInfo, error) {
	panic("unimplemented")
}

// Insert implements domain.IUserRepository.
func (u *Mock) Insert(ctx context.Context, usr *domain.User) error {
	panic("unimplemented")
}

// OAuth2SaveInfo implements domain.IUserRepository.
func (u *Mock) OAuth2SaveInfo(ctx context.Context, user *domain.User) error {
	panic("unimplemented")
}

// SaveInfo implements domain.IUserRepository.
func (u *Mock) SaveInfo(ctx context.Context, user *domain.User) error {
	panic("unimplemented")
}

// TransactionSaveOAuth2 implements domain.IUserRepository.
func (u *Mock) TransactionSaveOAuth2(ctx context.Context, data *domain.User) error {
	panic("unimplemented")
}

// TransactionSignUp implements domain.IUserRepository.
func (u *Mock) TransactionSignUp(ctx context.Context, user *domain.User, password string) error {
	panic("unimplemented")
}

// UpdateProperties implements domain.IUserRepository.
func (u *Mock) UpdateProperties(ctx context.Context, query string, user *domain.User) error {
	panic("unimplemented")
}

// Use implements domain.IUserRepository.
func (u *Mock) Use(tx *gorm.DB) IUserRepository {
	panic("unimplemented")
}
