// Package users
// Author: Duc Hung Ho @kieranhoo
// Description: users repository implementation
package users

import (
	"context"
	"errors"
	"fmt"
	"swclabs/swipecore/internal/core/repository/accounts"
	"swclabs/swipecore/pkg/lib/jwt"

	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"
	"swclabs/swipecore/pkg/utils"

	"gorm.io/gorm"
)

type Users struct {
	conn *gorm.DB
}

func New(conn *gorm.DB) *Users {
	return &Users{
		conn: conn,
	}
}

// GetByEmail implements domain.IUserRepository.
func (usr *Users) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := usr.conn.WithContext(ctx).
		Table(domain.UsersTable).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Insert implements domain.IUserRepository.
func (usr *Users) Insert(ctx context.Context, _usr *domain.User) error {
	return db.SafeWriteQuery(
		ctx,
		usr.conn,
		InsertIntoUsers,
		_usr.Email, _usr.PhoneNumber, _usr.FirstName, _usr.LastName, _usr.Image,
	)
}

// Info implements domain.IUserRepository.
func (usr *Users) Info(ctx context.Context, email string) (*domain.UserInfo, error) {
	data := new(domain.UserInfo)
	if err := usr.conn.WithContext(ctx).Raw(SelectUserInfo, email).
		Scan(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// SaveInfo implements domain.IUserRepository.
func (usr *Users) SaveInfo(ctx context.Context, user *domain.User) error {
	if user.Email == "" {
		return errors.New("missing key: email ")
	}
	if user.FirstName != "" {
		if err := db.SafeWriteQuery(
			ctx, usr.conn, UpdateUsersFirstname, user.FirstName, user.Email,
		); err != nil {
			return err
		}
	}
	if user.LastName != "" {
		if err := db.SafeWriteQuery(
			ctx, usr.conn, UpdateUsersFirstname, user.FirstName, user.Email,
		); err != nil {
			return err
		}
	}
	if user.Image != "" {
		if err := db.SafeWriteQuery(
			ctx, usr.conn, UpdateUsersImage, user.Image, user.Email,
		); err != nil {
			return err
		}
	}
	if user.PhoneNumber != "" {
		if err := db.SafeWriteQuery(
			ctx, usr.conn, UpdateUsersPhoneNumber, user.PhoneNumber, user.Email,
		); err != nil {
			return err
		}
	}
	return nil
}

// UpdateProperties implements domain.IUserRepository.
func (usr *Users) UpdateProperties(
	ctx context.Context, query string, user *domain.User) error {
	switch query {
	case UpdateUsersLastname:
		if err := db.SafeWriteQuery(
			ctx, usr.conn,
			UpdateUsersLastname, user.LastName, user.Email,
		); err != nil {
			return err
		}
	case UpdateUsersFirstname:
		if err := db.SafeWriteQuery(
			ctx, usr.conn,
			UpdateUsersFirstname, user.FirstName, user.Email,
		); err != nil {
			return err
		}
	case UpdateUsersPhoneNumber:
		if err := db.SafeWriteQuery(
			ctx, usr.conn,
			UpdateUsersPhoneNumber, user.PhoneNumber, user.Email,
		); err != nil {
			return err
		}
	case UpdateUsersImage:
		if err := db.SafeWriteQuery(
			ctx, usr.conn,
			UpdateUsersImage, user.Image, user.Email,
		); err != nil {
			return err
		}
	}
	return errors.New("unknown :" + query)
}

// OAuth2SaveInfo implements domain.IUserRepository.
func (usr *Users) OAuth2SaveInfo(ctx context.Context, user *domain.User) error {
	return db.SafeWriteQuery(
		ctx, usr.conn, InsertUsersConflict, user.Email, user.PhoneNumber,
		user.FirstName, user.LastName, user.Image,
	)
}

// TransactionSignUp implements domain.IUserRepository.
func (usr *Users) TransactionSignUp(
	ctx context.Context, user *domain.User, password string) error {
	return usr.conn.Transaction(func(tx *gorm.DB) error {
		hash, err := jwt.GenPassword(password)
		if err != nil {
			return err
		}
		if err := New(tx).Insert(ctx, user); err != nil {
			return err
		}
		userInfo, err := New(tx).GetByEmail(ctx, user.Email)
		if err != nil {
			return err
		}
		return accounts.New(tx).Insert(ctx, &domain.Account{
			Username: fmt.Sprintf("user#%d", userInfo.Id),
			Password: hash,
			Role:     "Customer",
			Email:    user.Email,
			Type:     "swc",
		})
	})
}

// TransactionSaveOAuth2 implements domain.IUserRepository.
func (usr *Users) TransactionSaveOAuth2(ctx context.Context, user *domain.User) error {
	return usr.conn.Transaction(func(tx *gorm.DB) error {
		hash, err := jwt.GenPassword(utils.RandomString(18))
		if err != nil {
			return err
		}
		if err := New(tx).OAuth2SaveInfo(ctx, user); err != nil {
			return err
		}
		userInfo, err := New(tx).GetByEmail(ctx, user.Email)
		if err != nil {
			return err
		}
		return accounts.New(tx).Insert(ctx, &domain.Account{
			Username: fmt.Sprintf("user#%d", userInfo.Id),
			Password: hash,
			Role:     "Customer",
			Email:    user.Email,
			Type:     "oauth2-google",
		})
	})
}

// GetByPhone implements domain.IUserRepository.
func (usr *Users) GetByPhone(
	ctx context.Context, nPhone string) (*domain.User, error) {
	var user domain.User
	if err := usr.conn.WithContext(ctx).
		Table(domain.UsersTable).
		Where("phone_number = ?", nPhone).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
