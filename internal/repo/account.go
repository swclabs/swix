package repo

import (
	"example/swiftcart/internal/model"
	"example/swiftcart/pkg/db"
	"time"

	"gorm.io/gorm"
)

type Accounts struct {
	data *model.Account
	conn *gorm.DB
}

func NewAccounts() IAccounts {
	_conn, err := db.Connection()
	if err != nil {
		panic(err)
	}
	return &Accounts{
		conn: _conn,
		data: &model.Account{},
	}
}

func (account *Accounts) GetByEmail(email string) (*model.Account, error) {
	if err := account.conn.Table("accounts").Where("email = ?", email).First(account.data).Error; err != nil {
		return nil, err
	}
	return account.data, nil
}

func (account *Accounts) Insert(acc *model.Account) error {
	createdAt := time.Now().UTC().Format(time.RFC3339)
	return account.conn.Exec(
		"INSERT INTO accounts (username, role, email, password, created_at) VALUES (?, ?, ?, ?, ?)",
		acc.Username, acc.Role, acc.Email, acc.Password, createdAt,
	).Error
}

func (account *Accounts) ChangePassword(pw string) error {
	return nil
}
