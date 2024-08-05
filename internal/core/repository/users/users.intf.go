package users

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/domain/model"
)

// IUserRepository Users Repository interface
// implement at /internal/repository/users.go
type IUserRepository interface {

	// GetByEmail retrieves a Users based on email address.
	GetByEmail(ctx context.Context, email string) (*entity.Users, error)

	// GetByPhone retrieves a Users based on phone number.
	GetByPhone(ctx context.Context, nPhone string) (*entity.Users, error)

	// GetByID retrieves a Users based on ID.
	GetByID(ctx context.Context, id int64) (*entity.Users, error)

	// Insert inserts a new Users into the database.
	Insert(ctx context.Context, usr entity.Users) error

	// Info retrieves Users information based on email address.
	Info(ctx context.Context, email string) (*model.Users, error)

	// Save saves Users information.
	Save(ctx context.Context, user entity.Users) error

	// OAuth2SaveInfo saves Users information from OAuth2 login.
	OAuth2SaveInfo(ctx context.Context, user entity.Users) error
}
