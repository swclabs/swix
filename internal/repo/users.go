package repo

import (
	"example/swiftcart/internal/model"
	"example/swiftcart/internal/schema"
	"example/swiftcart/pkg/db"

	"gorm.io/gorm"
)

type Users struct {
	conn *gorm.DB
	data *model.User
}

func NewUsers() IUsers {
	_conn, err := db.Connection()
	if err != nil {
		panic(err)
	}
	return &Users{
		conn: _conn,
		data: &model.User{},
	}
}

func (usr *Users) GetByEmail(email string) (*model.User, error) {
	if err := usr.conn.Table("users").Where("email = ?", email).First(usr.data).Error; err != nil {
		return nil, err
	}
	return usr.data, nil
}

func (usr *Users) Insert(_usr *model.User) error {
	return usr.conn.Exec(
		`INSERT INTO users (email, phone_number, first_name, last_name, image) VALUES (?,?,?,?,?)`,
		_usr.Email, _usr.PhoneNumber, _usr.FirstName, _usr.LastName, _usr.Image,
	).Error
}

func (usr *Users) Infor(email string) (*schema.InforResponse, error) {
	data := new(schema.InforResponse)
	if err := usr.conn.Raw(`
		SELECT users.email, phone_number, first_name, last_name, image, username, role
		FROM users JOIN accounts ON users.email = accounts.email
		WHERE users.email = ?;
	`, email).Scan(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
