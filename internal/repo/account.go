// Author: Duc Hung Ho @ikierans
// Description: account repository implementation
package repo

import (
	"errors"
	"time"

	"swclabs/swiftcart/internal/model"
	"swclabs/swiftcart/pkg/db"
	"swclabs/swiftcart/pkg/db/queries"

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
		queries.InsertIntoAccounts,
		acc.Username, acc.Role, acc.Email, acc.Password, createdAt, acc.Type,
	).Error
}

func (account *Accounts) SaveInfo(acc *model.Account) error {
	if acc.Email == "" {
		return errors.New("missing key: email ")
	}
	if acc.Username != "" {
		if err := account.conn.Exec(
			queries.UpdateAccountsUsername,
			acc.Username, acc.Email,
		).Error; err != nil {
			return err
		}
	}
	if acc.Password != "" {
		if err := account.conn.Exec(
			queries.UpdateAccountsPassword,
			acc.Password, acc.Email,
		).Error; err != nil {
			return err
		}
	}
	if acc.Role != "" {
		if err := account.conn.Exec(
			queries.UpdateAccountsRole,
			acc.Role, acc.Email,
		).Error; err != nil {
			return err
		}
	}
	return nil
}
