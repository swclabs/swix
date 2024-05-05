// Package repository
// Author: Duc Hung Ho @kieranhoo
// Description: account repository implementation
package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"
	"swclabs/swipecore/pkg/db/queries"

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

// Use implements domain.IAccountRepository.
func (account *Accounts) Use(tx *gorm.DB) domain.IAccountRepository {
	account.conn = tx
	return account
}

// GetByEmail implements domain.IAccountRepository.
func (account *Accounts) GetByEmail(ctx context.Context, email string) (*domain.Account, error) {
	if err := account.conn.WithContext(ctx).Table("accounts").Where("email = ?", email).First(account.data).Error; err != nil {
		return nil, err
	}
	return account.data, nil
}

// Insert implements domain.IAccountRepository.
func (account *Accounts) Insert(ctx context.Context, acc *domain.Account) error {
	createdAt := time.Now().UTC().Format(time.RFC3339)
	return db.SafeWriteQuery(
		ctx,
		account.conn,
		queries.InsertIntoAccounts,
		acc.Username, acc.Role, acc.Email, acc.Password, createdAt, acc.Type,
	)
}

// SaveInfo implements domain.IAccountRepository.
func (account *Accounts) SaveInfo(ctx context.Context, acc *domain.Account) error {
	if acc.Email == "" {
		return errors.New("missing key: email ")
	}
	if acc.Username != "" {
		if err := db.SafeWriteQuery(
			ctx,
			account.conn,
			queries.UpdateAccountsUsername,
			acc.Username, acc.Email,
		); err != nil {
			return err
		}

	}
	if acc.Password != "" {
		if err := db.SafeWriteQuery(
			ctx,
			account.conn,
			queries.UpdateAccountsPassword,
			acc.Password, acc.Email,
		); err != nil {
			return err
		}
	}
	if acc.Role != "" {
		if err := db.SafeWriteQuery(
			ctx,
			account.conn,
			queries.UpdateAccountsRole,
			acc.Role, acc.Email,
		); err != nil {
			return err
		}
	}
	return nil
}
