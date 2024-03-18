// Package repo
// Author: Duc Hung Ho @kieranhoo
// Description: users repository implementation
package repo

import (
	"errors"
	"fmt"
	"log"

	"github.com/swclabs/swipe-api/internal/domain"
	"github.com/swclabs/swipe-api/pkg/db"
	"github.com/swclabs/swipe-api/pkg/db/queries"
	"github.com/swclabs/swipe-api/pkg/tools"
	"github.com/swclabs/swipe-api/pkg/utils"

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

func (usr *Users) GetByEmail(email string) (*domain.User, error) {
	if err := usr.conn.Table("users").Where("email = ?", email).First(usr.data).Error; err != nil {
		return nil, err
	}
	return usr.data, nil
}

func (usr *Users) Insert(_usr *domain.User) error {
	// return usr.conn.Exec(
	// 	queries.InsertIntoUsers,
	// 	_usr.Email, _usr.PhoneNumber, _usr.FirstName, _usr.LastName, _usr.Image,
	// ).Error

	return db.SafeWriteQuery(
		usr.conn,
		queries.InsertIntoUsers,
		_usr.Email, _usr.PhoneNumber, _usr.FirstName, _usr.LastName, _usr.Image,
	)
}

func (usr *Users) Info(email string) (*domain.UserInfo, error) {
	data := new(domain.UserInfo)
	if err := usr.conn.Raw(queries.SelectUserInfo, email).Scan(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (usr *Users) SaveInfo(user *domain.User) error {
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
			usr.conn,
			queries.UpdateUsersPhoneNumber,
			user.PhoneNumber, user.Email,
		); err != nil {
			return err
		}
	}
	return nil
}

func (usr *Users) OAuth2SaveInfo(user *domain.User) error {
	// return usr.conn.Exec(
	// 	queries.InsertUsersConflict,
	// 	user.Email,
	// 	user.PhoneNumber,
	// 	user.FirstName,
	// 	user.LastName,
	// 	user.Image,
	// ).Error

	return db.SafeWriteQuery(
		usr.conn,
		queries.InsertUsersConflict,
		user.Email,
		user.PhoneNumber,
		user.FirstName,
		user.LastName,
		user.Image,
	)
}

func (usr *Users) SignUp(user *domain.User, password string) error {
	account := NewAccounts()
	return usr.conn.Transaction(func(tx *gorm.DB) error {
		hash, err := tools.GenPassword(password)
		if err != nil {
			return err
		}
		if err := usr.Insert(user); err != nil {
			return err
		}
		userInfo, err := usr.GetByEmail(user.Email)
		if err != nil {
			return err
		}
		return account.Insert(&domain.Account{
			Username: fmt.Sprintf("user#%d", userInfo.UserID),
			Password: hash,
			Role:     "Customer",
			Email:    user.Email,
			Type:     "swc",
		})
	})
}

func (usr *Users) SaveOAuth2(user *domain.User) error {
	account := NewAccounts()
	return usr.conn.Transaction(func(tx *gorm.DB) error {
		hash, err := tools.GenPassword(utils.RandomString(18))
		if err != nil {
			return err
		}
		if err := usr.OAuth2SaveInfo(user); err != nil {
			return err
		}
		userInfo, err := usr.GetByEmail(user.Email)
		if err != nil {
			return err
		}
		return account.Insert(&domain.Account{
			Username: fmt.Sprintf("user#%d", userInfo.UserID),
			Password: hash,
			Role:     "Customer",
			Email:    user.Email,
			Type:     "oauth2-google",
		})
	})
}
