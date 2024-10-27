package users

import (
	"context"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/internal/core/domain/model"
)

// IUsers Users Repository interface
// implement at /internal/repos/users.go
type IUsers interface {

	// GetByEmail retrieves a Users based on email address.
	GetByEmail(ctx context.Context, email string) (*entity.User, error)

	// GetByPhone retrieves a Users based on phone number.
	GetByPhone(ctx context.Context, nPhone string) (*entity.User, error)

	// GetByID retrieves a Users based on ID.
	GetByID(ctx context.Context, id int64) (*entity.User, error)

	// Insert inserts a new Users into the database.
	Insert(ctx context.Context, usr entity.User) (int64, error)

	// Info retrieves Users information based on email address.
	Info(ctx context.Context, email string) (*model.Users, error)

	// Save saves Users information.
	Save(ctx context.Context, user entity.User) error

	// OAuth2SaveInfo saves Users information from OAuth2 login.
	OAuth2SaveInfo(ctx context.Context, user entity.User) error
}
