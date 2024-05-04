package domain

import (
	"context"

	"gorm.io/gorm"
)

// IUserRepository User Repository interface
// implement at /internal/repository/user.go
type IUserRepository interface {
	Use(tx *gorm.DB) IUserRepository

	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByPhone(ctx context.Context, nPhone string) (*User, error)
	Insert(ctx context.Context, usr *User) error
	Info(ctx context.Context, email string) (*UserInfo, error)
	SaveInfo(ctx context.Context, user *User) error
	OAuth2SaveInfo(ctx context.Context, user *User) error
	TransactionSignUp(ctx context.Context, user *User, password string) error
	TransactionSaveOAuth2(ctx context.Context, data *User) error
	UpdateProperties(ctx context.Context, query string, user *User) error
}
