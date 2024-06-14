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
	conn *pgx.Conn
}

func New(conn *pgx.Conn) *Users {
	return &Users{conn}
}

var _ IUserRepository = (*Users)(nil)

// GetByEmail implements domain.IUserRepository.
func (usr *Users) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	rows, err := usr.conn.Query(ctx, SelectByEmail, email)
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
	return db.SafePgxWriteQuery(
		ctx, usr.conn,
		InsertIntoUsers,
		_usr.Email, _usr.PhoneNumber, _usr.FirstName, _usr.LastName, _usr.Image,
	)
}

// Info implements domain.IUserRepository.
func (usr *Users) Info(ctx context.Context, email string) (*domain.UserInfo, error) {
	rows, err := usr.conn.Query(ctx, SelectUserInfo, email)
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
		if err := db.SafePgxWriteQuery(
			ctx, usr.conn, UpdateUsersFirstname, user.FirstName, user.Email,
		); err != nil {
			return err
		}
	}
	if user.LastName != "" {
		if err := db.SafePgxWriteQuery(
			ctx, usr.conn, UpdateUsersFirstname, user.FirstName, user.Email,
		); err != nil {
			return err
		}
	}
	if user.Image != "" {
		if err := db.SafePgxWriteQuery(
			ctx, usr.conn, UpdateUsersImage, user.Image, user.Email,
		); err != nil {
			return err
		}
	}
	if user.PhoneNumber != "" {
		if err := db.SafePgxWriteQuery(
			ctx, usr.conn, UpdateUsersPhoneNumber, user.PhoneNumber, user.Email,
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
		if err := db.SafePgxWriteQuery(
			ctx, usr.conn,
			UpdateUsersLastname, user.LastName, user.Email,
		); err != nil {
			return err
		}
	case UpdateUsersFirstname:
		if err := db.SafePgxWriteQuery(
			ctx, usr.conn,
			UpdateUsersFirstname, user.FirstName, user.Email,
		); err != nil {
			return err
		}
	case UpdateUsersPhoneNumber:
		if err := db.SafePgxWriteQuery(
			ctx, usr.conn,
			UpdateUsersPhoneNumber, user.PhoneNumber, user.Email,
		); err != nil {
			return err
		}
	case UpdateUsersImage:
		if err := db.SafePgxWriteQuery(
			ctx, usr.conn,
			UpdateUsersImage, user.Image, user.Email,
		); err != nil {
			return err
		}
	}
	return errors.New("unknown :" + query)
}

// OAuth2SaveInfo implements domain.IUserRepository.
func (usr *Users) OAuth2SaveInfo(ctx context.Context, user domain.User) error {
	return db.SafePgxWriteQuery(
		ctx, usr.conn, InsertUsersConflict, user.Email, user.PhoneNumber,
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
	tx, err := usr.conn.Begin(ctx)
	if err != nil {
		return err
	}
	if err := New(tx.Conn()).Insert(ctx, user); err != nil {
		tx.Rollback(ctx)
		return err
	}

	userInfo, err := New(tx.Conn()).GetByEmail(ctx, user.Email)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	if err := accounts.New(tx.Conn()).Insert(ctx, domain.Account{
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
	tx, err := usr.conn.Begin(ctx)
	if err != nil {
		return err
	}
	if err := New(tx.Conn()).OAuth2SaveInfo(ctx, user); err != nil {
		tx.Rollback(ctx)
		return err
	}
	userInfo, err := New(tx.Conn()).GetByEmail(ctx, user.Email)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	if err := accounts.New(tx.Conn()).Insert(ctx, domain.Account{
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
	rows, err := usr.conn.Query(ctx, selectByPhone, nPhone)
	if err != nil {
		return nil, err
	}
	user, err := pgx.CollectOneRow[domain.User](rows, pgx.RowToStructByName[domain.User])
	if err != nil {
		return nil, err
	}
	return &user, nil
}
