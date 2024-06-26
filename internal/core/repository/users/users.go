// Package users
// Author: Duc Hung Ho @kyeranyo
// Description: users repository implementation
package users

import (
	"context"
	"errors"
	"fmt"
	"swclabs/swipecore/internal/core/repository/accounts"
	"swclabs/swipecore/pkg/lib/jwt"
	"swclabs/swipecore/pkg/utils"

	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"

	"github.com/jackc/pgx/v5"
)

type Users struct {
	db db.IDatabase
}

func New(conn db.IDatabase) IUserRepository {
	return &Users{conn}
}

var _ IUserRepository = (*Users)(nil)

// GetByEmail implements domain.IUserRepository.
func (usr *Users) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	rows, err := usr.db.Query(ctx, SelectByEmail, email)
	if err != nil {
		return nil, err
	}
	user, err := pgx.CollectOneRow[domain.User](rows, pgx.RowToStructByName[domain.User])
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Insert implements domain.IUserRepository.
func (usr *Users) Insert(ctx context.Context, _usr domain.User) error {
	return usr.db.SafeWrite(
		ctx,
		InsertIntoUsers,
		_usr.Email, _usr.PhoneNumber, _usr.FirstName, _usr.LastName, _usr.Image,
	)
}

// Info implements domain.IUserRepository.
func (usr *Users) Info(ctx context.Context, email string) (*domain.UserInfo, error) {
	rows, err := usr.db.Query(ctx, SelectUserInfo, email)
	if err != nil {
		return nil, err
	}
	user, err := pgx.CollectOneRow[domain.UserInfo](rows, pgx.RowToStructByName[domain.UserInfo])
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// SaveInfo implements domain.IUserRepository.
func (usr *Users) SaveInfo(ctx context.Context, user domain.User) error {
	if user.Email == "" {
		return errors.New("missing key: email ")
	}
	if user.FirstName != "" {
		if err := usr.db.SafeWrite(
			ctx, UpdateUsersFirstname, user.FirstName, user.Email,
		); err != nil {
			return err
		}
	}
	if user.LastName != "" {
		if err := usr.db.SafeWrite(
			ctx, UpdateUsersFirstname, user.FirstName, user.Email,
		); err != nil {
			return err
		}
	}
	if user.Image != "" {
		if err := usr.db.SafeWrite(
			ctx, UpdateUsersImage, user.Image, user.Email,
		); err != nil {
			return err
		}
	}
	if user.PhoneNumber != "" {
		if err := usr.db.SafeWrite(
			ctx, UpdateUsersPhoneNumber, user.PhoneNumber, user.Email,
		); err != nil {
			return err
		}
	}
	return nil
}

// UpdateProperties implements domain.IUserRepository.
func (usr *Users) UpdateProperties(
	ctx context.Context, query string, user domain.User) error {
	switch query {
	case UpdateUsersLastname:
		if err := usr.db.SafeWrite(ctx,
			UpdateUsersLastname, user.LastName, user.Email,
		); err != nil {
			return err
		}
	case UpdateUsersFirstname:
		if err := usr.db.SafeWrite(ctx,
			UpdateUsersFirstname, user.FirstName, user.Email,
		); err != nil {
			return err
		}
	case UpdateUsersPhoneNumber:
		if err := usr.db.SafeWrite(ctx,
			UpdateUsersPhoneNumber, user.PhoneNumber, user.Email,
		); err != nil {
			return err
		}
	case UpdateUsersImage:
		if err := usr.db.SafeWrite(ctx,
			UpdateUsersImage, user.Image, user.Email,
		); err != nil {
			return err
		}
	}
	return errors.New("unknown :" + query)
}

// OAuth2SaveInfo implements domain.IUserRepository.
func (usr *Users) OAuth2SaveInfo(ctx context.Context, user domain.User) error {
	return usr.db.SafeWrite(
		ctx, InsertUsersConflict, user.Email, user.PhoneNumber,
		user.FirstName, user.LastName, user.Image,
	)
}

// TransactionSignUp implements domain.IUserRepository.
func (usr *Users) TransactionSignUp(
	ctx context.Context, user domain.User, password string) error {
	hash, err := jwt.GenPassword(password)
	if err != nil {
		return err
	}
	tx, err := db.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	if err := New(tx).Insert(ctx, user); err != nil {
		tx.Rollback(ctx)
		return err
	}

	userInfo, err := New(tx).GetByEmail(ctx, user.Email)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	if err := accounts.New(tx).Insert(ctx, domain.Account{
		Username: fmt.Sprintf("user#%d", userInfo.Id),
		Password: hash,
		Role:     "Customer",
		Email:    user.Email,
		Type:     "swc",
	}); err != nil {
		tx.Rollback(ctx)
		return err
	}
	return tx.Commit(ctx)
}

// TransactionSaveOAuth2 implements domain.IUserRepository.
func (usr *Users) TransactionSaveOAuth2(ctx context.Context, user domain.User) error {
	hash, err := jwt.GenPassword(utils.RandomString(18))
	if err != nil {
		return err
	}
	tx, err := db.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	if err := New(tx).OAuth2SaveInfo(ctx, user); err != nil {
		tx.Rollback(ctx)
		return err
	}
	userInfo, err := New(tx).GetByEmail(ctx, user.Email)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	if err := accounts.New(tx).Insert(ctx, domain.Account{
		Username: fmt.Sprintf("user#%d", userInfo.Id),
		Password: hash,
		Role:     "Customer",
		Email:    user.Email,
		Type:     "oauth2-google",
	}); err != nil {
		tx.Rollback(ctx)
		return err
	}
	return tx.Commit(ctx)
}

// GetByPhone implements domain.IUserRepository.
func (usr *Users) GetByPhone(ctx context.Context, nPhone string) (*domain.User, error) {
	rows, err := usr.db.Query(ctx, selectByPhone, nPhone)
	if err != nil {
		return nil, err
	}
	user, err := db.CollectOneRow[domain.User](rows)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
