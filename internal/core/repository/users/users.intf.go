package users

import (
	"context"
	"swclabs/swipecore/internal/core/domain"

	"gorm.io/gorm"
)

// IUserRepository User Repository interface
// implement at /internal/repository/user.go
type IUserRepository interface {
	Use(tx *gorm.DB) IUserRepository

	// GetByEmail retrieves a User based on email address.
	GetByEmail(ctx context.Context, email string) (*domain.User, error)

	// GetByPhone retrieves a User based on phone number.
	GetByPhone(ctx context.Context, nPhone string) (*domain.User, error)

	// Insert inserts a new User into the database.
	Insert(ctx context.Context, usr *domain.User) error

	// Info retrieves User information based on email address.
	Info(ctx context.Context, email string) (*domain.UserInfo, error)

	// SaveInfo saves User information.
	SaveInfo(ctx context.Context, user *domain.User) error

	// OAuth2SaveInfo saves User information from OAuth2 login.
	OAuth2SaveInfo(ctx context.Context, user *domain.User) error

	// TransactionSignUp signs up a User within a transaction.
	TransactionSignUp(ctx context.Context, user *domain.User, password string) error

	// TransactionSaveOAuth2 saves User information from OAuth2 login within a transaction.
	TransactionSaveOAuth2(ctx context.Context, data *domain.User) error

	// UpdateProperties updates User properties.
	UpdateProperties(ctx context.Context, query string, user *domain.User) error
}
