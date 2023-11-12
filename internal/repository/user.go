package repository

import (
	"example/komposervice/internal/model"
	"example/komposervice/pkg/db"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	conn *gorm.DB
	data *model.Users
}

func NewUsers() IUsers {
	_conn, err := db.Connection()
	if err != nil {
		panic(err)
	}
	return &Users{
		conn: _conn,
		data: &model.Users{},
	}
}

func (usr *Users) GetByEmail(email string) (*model.Users, error) {
	if err := usr.conn.Table("users").Where("email = ?", email).First(usr.data).Error; err != nil {
		return nil, err
	}
	return usr.data, nil
}

func (usr *Users) Create(_usr *model.Users) error {
	createAt := time.Now().UTC().Format(time.RFC3339)
	return usr.conn.Exec(
		`INSERT INTO users (username,hashed_password,full_name,email,password_changed_at,created_at)
		VALUES (?,?,?,?,?,?)`,
		_usr.Username, _usr.HashedPassword, _usr.FullName, _usr.Email, createAt, createAt,
	).Error
}

func (usr *Users) Empty() bool {
	return usr.data.Email == ""
}
