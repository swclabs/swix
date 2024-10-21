// Package users users repos implementation
package users

import (
	"context"
	"swclabs/swix/app"

	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/domain/model"
	"swclabs/swix/pkg/infra/cache"
	"swclabs/swix/pkg/infra/db"
)

// New creates a new instance of IUserRepository.
func New(conn db.IDatabase) IUsers {
	return &Users{conn}
}

var _ = app.Repos(Init)

// Init initializes the Users object with database and redis connection
func Init(conn db.IDatabase, cache cache.ICache) IUsers {
	return useCache(cache, New(conn))
}

var _ IUsers = (*Users)(nil)

// Users repos implementation
type Users struct {
	db db.IDatabase
}

// GetByID implements IUserRepository.
func (usr *Users) GetByID(ctx context.Context, id int64) (*entity.Users, error) {
	rows, err := usr.db.Query(ctx, selectByID, id)
	if err != nil {
		return nil, err
	}
	user, err := db.CollectRow[entity.Users](rows)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail implements IUserRepository.
func (usr *Users) GetByEmail(ctx context.Context, email string) (*entity.Users, error) {
	rows, err := usr.db.Query(ctx, selectByEmail, email)
	if err != nil {
		return nil, err
	}
	user, err := db.CollectRow[entity.Users](rows)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Insert implements IUserRepository.
func (usr *Users) Insert(ctx context.Context, _usr entity.Users) (int64, error) {
	return usr.db.SafeWriteReturn(
		ctx,
		insertIntoUsers,
		_usr.Email, _usr.PhoneNumber, _usr.FirstName, _usr.LastName, _usr.Image,
	)
}

// Info implements IUserRepository.
func (usr *Users) Info(ctx context.Context, email string) (*model.Users, error) {
	rows, err := usr.db.Query(ctx, selectUserInfo, email)
	if err != nil {
		return nil, err
	}
	user, err := db.CollectRow[model.Users](rows)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Save implements IUserRepository.
func (usr *Users) Save(ctx context.Context, user entity.Users) error {
	return usr.db.SafeWrite(ctx, updateInfo,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Image,
		user.PhoneNumber,
	)
}

// OAuth2SaveInfo implements IUserRepository.
func (usr *Users) OAuth2SaveInfo(ctx context.Context, user entity.Users) error {
	return usr.db.SafeWrite(
		ctx, insertUsersConflict, user.Email, user.PhoneNumber,
		user.FirstName, user.LastName, user.Image,
	)
}

// GetByPhone implements IUserRepository.
func (usr *Users) GetByPhone(ctx context.Context, nPhone string) (*entity.Users, error) {
	rows, err := usr.db.Query(ctx, selectByPhone, nPhone)
	if err != nil {
		return nil, err
	}
	user, err := db.CollectRow[entity.Users](rows)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
