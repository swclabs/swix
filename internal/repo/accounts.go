// Package repo
// Author: Duc Hung Ho @kieranhoo
// Description: account repository implementation
package repo

import (
	"errors"
	"log"
	"time"

	"github.com/swclabs/swipe-api/internal/domain"
	"github.com/swclabs/swipe-api/pkg/db"
	"github.com/swclabs/swipe-api/pkg/db/queries"

	"gorm.io/gorm"
)

type Accounts struct {
	data *domain.Account
	conn *gorm.DB
}

func NewAccounts() domain.IAccountRepository {
	_conn, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	return &Accounts{
		conn: _conn,
		data: &domain.Account{},
	}
}

func (account *Accounts) GetByEmail(email string) (*domain.Account, error) {
	if err := account.conn.Table("accounts").Where("email = ?", email).First(account.data).Error; err != nil {
		return nil, err
	}
	return account.data, nil
}

func (account *Accounts) Insert(acc *domain.Account) error {
	createdAt := time.Now().UTC().Format(time.RFC3339)
	// return account.conn.Exec(
	// 	queries.InsertIntoAccounts,
	// 	acc.Username, acc.Role, acc.Email, acc.Password, createdAt, acc.Type,
	// ).Error
	return db.SafeWriteQuery(
		account.conn,
		queries.InsertIntoAccounts,
		acc.Username, acc.Role, acc.Email, acc.Password, createdAt, acc.Type,
	)
}

func (account *Accounts) SaveInfo(acc *domain.Account) error {
	if acc.Email == "" {
		return errors.New("missing key: email ")
	}
	if acc.Username != "" {
		// if err := account.conn.Exec(
		// 	queries.UpdateAccountsUsername,
		// 	acc.Username, acc.Email,
		// ).Error; err != nil {
		// 	return err
		// }

		if err := db.SafeWriteQuery(
			account.conn,
			queries.UpdateAccountsUsername,
			acc.Username, acc.Email,
		); err != nil {
			return err
		}

	}
	if acc.Password != "" {
		// if err := account.conn.Exec(
		// 	queries.UpdateAccountsPassword,
		// 	acc.Password, acc.Email,
		// ).Error; err != nil {
		// 	return err
		// }

		if err := db.SafeWriteQuery(
			account.conn,
			queries.UpdateAccountsPassword,
			acc.Password, acc.Email,
		); err != nil {
			return err
		}
	}
	if acc.Role != "" {
		// if err := account.conn.Exec(
		// 	queries.UpdateAccountsRole,
		// 	acc.Role, acc.Email,
		// ).Error; err != nil {
		// 	return err
		// }

		if err := db.SafeWriteQuery(
			account.conn,
			queries.UpdateAccountsRole,
			acc.Role, acc.Email,
		); err != nil {
			return err
		}
	}
	return nil
}
