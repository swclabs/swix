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
)

type Users struct {
	db db.IDatabase
}

func New(conn db.IDatabase) IUserRepository {
	return useCache(&Users{conn})
}

var _ IUserRepository = (*Users)(nil)

// GetByEmail implements IUserRepository.
func (usr *Users) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	rows, err := usr.db.Query(ctx, selectByEmail, email)
	if err != nil {
		return nil, err
	}
	user, err := db.CollectOneRow[domain.User](rows)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Insert implements IUserRepository.
func (usr *Users) Insert(ctx context.Context, _usr domain.User) error {
	return usr.db.SafeWrite(
		ctx,
		insertIntoUsers,
		_usr.Email, _usr.PhoneNumber, _usr.FirstName, _usr.LastName, _usr.Image,
	)
}

// Info implements IUserRepository.
func (usr *Users) Info(ctx context.Context, email string) (*domain.UserInfo, error) {
	rows, err := usr.db.Query(ctx, selectUserInfo, email)
	if err != nil {
		return nil, err
	}
	user, err := db.CollectOneRow[domain.UserInfo](rows)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// SaveInfo implements IUserRepository.
func (usr *Users) SaveInfo(ctx context.Context, user domain.User) error {
	if user.Email == "" {
		return errors.New("missing key: email ")
	}
	if user.FirstName != "" {
		if err := usr.db.SafeWrite(
			ctx, updateUsersFirstname, user.FirstName, user.Email,
		); err != nil {
			return err
		}
	}
	if user.LastName != "" {
		if err := usr.db.SafeWrite(
			ctx, updateUsersFirstname, user.FirstName, user.Email,
		); err != nil {
			return err
		}
	}
	if user.Image != "" {
		if err := usr.db.SafeWrite(
			ctx, updateUsersImage, user.Image, user.Email,
		); err != nil {
			return err
		}
	}
	if user.PhoneNumber != "" {
		if err := usr.db.SafeWrite(
			ctx, updateUsersPhoneNumber, user.PhoneNumber, user.Email,
		); err != nil {
			return err
		}
	}
	return nil
}

// UpdateProperties implements IUserRepository.
func (usr *Users) UpdateProperties(
	ctx context.Context, query string, user domain.User) error {
	switch query {
	case updateUsersLastname:
		if err := usr.db.SafeWrite(ctx,
			updateUsersLastname, user.LastName, user.Email,
		); err != nil {
			return err
		}
	case updateUsersFirstname:
		if err := usr.db.SafeWrite(ctx,
			updateUsersFirstname, user.FirstName, user.Email,
		); err != nil {
			return err
		}
	case updateUsersPhoneNumber:
		if err := usr.db.SafeWrite(ctx,
			updateUsersPhoneNumber, user.PhoneNumber, user.Email,
		); err != nil {
			return err
		}
	case updateUsersImage:
		if err := usr.db.SafeWrite(ctx,
			updateUsersImage, user.Image, user.Email,
		); err != nil {
			return err
		}
	}
	return errors.New("unknown :" + query)
}

// OAuth2SaveInfo implements IUserRepository.
func (usr *Users) OAuth2SaveInfo(ctx context.Context, user domain.User) error {
	return usr.db.SafeWrite(
		ctx, insertUsersConflict, user.Email, user.PhoneNumber,
		user.FirstName, user.LastName, user.Image,
	)
}

// TransactionSignUp implements IUserRepository.
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

// TransactionSaveOAuth2 implements IUserRepository.
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

// GetByPhone implements IUserRepository.
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
