// Package repo
// Author: Duc Hung Ho @kieranhoo
// Description: users repository implementation
package repo

import (
	"context"
	"errors"
	"fmt"
	"log"

	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/pkg/db"
	"swclabs/swipe-api/pkg/db/queries"
	"swclabs/swipe-api/pkg/tools"
	"swclabs/swipe-api/pkg/utils"

	"gorm.io/gorm"
)

type Users struct {
	conn *gorm.DB
	data *domain.User
}

func NewUsers() domain.IUserRepository {
	_conn, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	return &Users{
		conn: _conn,
		data: &domain.User{},
	}
}

func (usr *Users) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	if err := usr.conn.WithContext(ctx).Table("users").Where("email = ?", email).First(usr.data).Error; err != nil {
		return nil, err
	}
	return usr.data, nil
}

func (usr *Users) Insert(ctx context.Context, _usr *domain.User) error {
	// return usr.conn.Exec(
	// 	queries.InsertIntoUsers,
	// 	_usr.Email, _usr.PhoneNumber, _usr.FirstName, _usr.LastName, _usr.Image,
	// ).Error

	return db.SafeWriteQuery(
		ctx,
		usr.conn,
		queries.InsertIntoUsers,
		_usr.Email, _usr.PhoneNumber, _usr.FirstName, _usr.LastName, _usr.Image,
	)
}

func (usr *Users) Info(ctx context.Context, email string) (*domain.UserInfo, error) {
	data := new(domain.UserInfo)
	if err := usr.conn.WithContext(ctx).Raw(queries.SelectUserInfo, email).Scan(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (usr *Users) SaveInfo(ctx context.Context, user *domain.User) error {
	if user.Email == "" {
		return errors.New("missing key: email ")
	}
	if user.FirstName != "" {
		// if err := usr.conn.Exec(
		// 	queries.UpdateUsersFirstname,
		// 	user.FirstName, user.Email,
		// ).Error; err != nil {
		// 	return err
		// }

		if err := db.SafeWriteQuery(
			ctx,
			usr.conn,
			queries.UpdateUsersFirstname,
			user.FirstName, user.Email,
		); err != nil {
			return err
		}
	}
	if user.LastName != "" {
		// if err := usr.conn.Exec(
		// 	queries.UpdateUsersLastname,
		// 	user.LastName, user.Email,
		// ).Error; err != nil {
		// 	return err
		// }

		if err := db.SafeWriteQuery(
			ctx,
			usr.conn,
			queries.UpdateUsersFirstname,
			user.FirstName, user.Email,
		); err != nil {
			return err
		}
	}
	if user.Image != "" {
		// if err := usr.conn.Exec(
		// 	queries.UpdateUsersImage,
		// 	user.Image, user.Email,
		// ).Error; err != nil {
		// 	return err
		// }

		if err := db.SafeWriteQuery(
			ctx,
			usr.conn,
			queries.UpdateUsersImage,
			user.Image, user.Email,
		); err != nil {
			return err
		}
	}
	if user.PhoneNumber != "" {
		// if err := usr.conn.Exec(
		// 	queries.UpdateUsersPhoneNumber,
		// 	user.PhoneNumber, user.Email,
		// ).Error; err != nil {
		// 	return err
		// }

		if err := db.SafeWriteQuery(
			ctx,
			usr.conn,
			queries.UpdateUsersPhoneNumber,
			user.PhoneNumber, user.Email,
		); err != nil {
			return err
		}
	}
	return nil
}

func (usr *Users) UpdateProperties(ctx context.Context, query string, user *domain.User) error {
	switch query {
	case queries.UpdateUsersLastname:
		if err := db.SafeWriteQuery(
			ctx,
			usr.conn,
			queries.UpdateUsersLastname,
			user.LastName, user.Email,
		); err != nil {
			return err
		}
	case queries.UpdateUsersFirstname:
		if err := db.SafeWriteQuery(
			ctx,
			usr.conn,
			queries.UpdateUsersFirstname,
			user.FirstName, user.Email,
		); err != nil {
			return err
		}
	case queries.UpdateUsersPhoneNumber:
		if err := db.SafeWriteQuery(
			ctx,
			usr.conn,
			queries.UpdateUsersPhoneNumber,
			user.PhoneNumber, user.Email,
		); err != nil {
			return err
		}
	case queries.UpdateUsersImage:
		if err := db.SafeWriteQuery(
			ctx,
			usr.conn,
			queries.UpdateUsersImage,
			user.Image, user.Email,
		); err != nil {
			return err
		}
	}
	return errors.New("unknown :" + query)
}

func (usr *Users) OAuth2SaveInfo(ctx context.Context, user *domain.User) error {
	// return usr.conn.Exec(
	// 	queries.InsertUsersConflict,
	// 	user.Email,
	// 	user.PhoneNumber,
	// 	user.FirstName,
	// 	user.LastName,
	// 	user.Image,
	// ).Error

	return db.SafeWriteQuery(
		ctx,
		usr.conn,
		queries.InsertUsersConflict,
		user.Email,
		user.PhoneNumber,
		user.FirstName,
		user.LastName,
		user.Image,
	)
}

func (usr *Users) TransactionSignUp(ctx context.Context, user *domain.User, password string) error {
	account := NewAccounts()
	return usr.conn.Transaction(func(tx *gorm.DB) error {
		hash, err := tools.GenPassword(password)
		if err != nil {
			return err
		}
		if err := usr.Insert(ctx, user); err != nil {
			return err
		}
		userInfo, err := usr.GetByEmail(ctx, user.Email)
		if err != nil {
			return err
		}
		return account.Insert(ctx, &domain.Account{
			Username: fmt.Sprintf("user#%d", userInfo.UserID),
			Password: hash,
			Role:     "Customer",
			Email:    user.Email,
			Type:     "swc",
		})
	})
}

func (usr *Users) TransactionSaveOAuth2(ctx context.Context, user *domain.User) error {
	account := NewAccounts()
	return usr.conn.Transaction(func(tx *gorm.DB) error {
		hash, err := tools.GenPassword(utils.RandomString(18))
		if err != nil {
			return err
		}
		if err := usr.OAuth2SaveInfo(ctx, user); err != nil {
			return err
		}
		userInfo, err := usr.GetByEmail(ctx, user.Email)
		if err != nil {
			return err
		}
		return account.Insert(ctx, &domain.Account{
			Username: fmt.Sprintf("user#%d", userInfo.UserID),
			Password: hash,
			Role:     "Customer",
			Email:    user.Email,
			Type:     "oauth2-google",
		})
	})
}
