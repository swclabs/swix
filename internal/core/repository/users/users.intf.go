package users

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

// IUserRepository Users Repository interface
// implement at /internal/repository/users.go
type IUserRepository interface {

	// GetByEmail retrieves a Users based on email address.
	GetByEmail(ctx context.Context, email string) (*domain.Users, error)

	// GetByPhone retrieves a Users based on phone number.
	GetByPhone(ctx context.Context, nPhone string) (*domain.Users, error)

	// Insert inserts a new Users into the database.
	Insert(ctx context.Context, usr domain.Users) error

	// Info retrieves Users information based on email address.
	Info(ctx context.Context, email string) (*domain.UserSchema, error)

	// SaveInfo saves Users information.
	SaveInfo(ctx context.Context, user domain.Users) error

	// OAuth2SaveInfo saves Users information from OAuth2 login.
	OAuth2SaveInfo(ctx context.Context, user domain.Users) error

	// UpdateProperties updates Users properties.
	UpdateProperties(ctx context.Context, query string, user domain.Users) error
}
