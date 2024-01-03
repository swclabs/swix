// Author: Duc Hung Ho @ikierans
// Description: users repository implementation
package repo

import (
	"errors"
	"swclabs/swiftcart/internal/domain"
	"swclabs/swiftcart/pkg/db"
	"swclabs/swiftcart/pkg/db/queries"

	"gorm.io/gorm"
)

type Users struct {
	conn *gorm.DB
	data *domain.User
}

func NewUsers() domain.IUserRepository {
	_conn, err := db.Connection()
	if err != nil {
		panic(err)
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
	return usr.conn.Exec(
		queries.InsertIntoUsers,
		_usr.Email, _usr.PhoneNumber, _usr.FirstName, _usr.LastName, _usr.Image,
	).Error
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
		if err := usr.conn.Exec(
			queries.UpdateUsersFirstname,
			user.FirstName, user.Email,
		).Error; err != nil {
			return err
		}
	}
	if user.LastName != "" {
		if err := usr.conn.Exec(
			queries.UpdateUsersLastname,
			user.LastName, user.Email,
		).Error; err != nil {
			return err
		}
	}
	if user.Image != "" {
		if err := usr.conn.Exec(
			queries.UpdateUsersImage,
			user.Image, user.Email,
		).Error; err != nil {
			return err
		}
	}
	if user.PhoneNumber != "" {
		if err := usr.conn.Exec(
			queries.UpdateUsersPhoneNumber,
			user.PhoneNumber, user.Email,
		).Error; err != nil {
			return err
		}
	}
	return nil
}

func (usr *Users) OAuth2SaveInfo(user *domain.User) error {
	return usr.conn.Exec(
		queries.InsertUsersConflict,
		user.Email,
		user.PhoneNumber,
		user.FirstName,
		user.LastName,
		user.Image,
	).Error
}
