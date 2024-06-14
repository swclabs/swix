// Package accounts
// Author: Duc Hung Ho @kyeranyo
// Description: account repository implementation
package accounts

import (
	"context"
	"errors"
	"time"

	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"

	"github.com/jackc/pgx/v5"
)

type Accounts struct {
	conn *pgx.Conn
}

func New(conn *pgx.Conn) *Accounts {
	return &Accounts{conn}
}

// GetByEmail implements domain.IAccountRepository.
func (account *Accounts) GetByEmail(
	ctx context.Context, email string) (*domain.Account, error) {
	rows, err := account.conn.Query(ctx, SelectByEmail, email)
	if err != nil {
		return nil, err
	}
	acc, err := pgx.CollectOneRow[domain.Account](rows, pgx.RowToStructByName[domain.Account])
	if err != nil {
		return nil, err
	}
	return &acc, nil
}

// Insert implements domain.IAccountRepository.
func (account *Accounts) Insert(
	ctx context.Context, acc domain.Account) error {
	createdAt := time.Now().UTC().Format(time.RFC3339)
	return db.SafePgxWriteQuery(
		ctx, account.conn,
		InsertIntoAccounts,
		acc.Username, acc.Role, acc.Email, acc.Password,
		createdAt, acc.Type,
	)
}

// SaveInfo implements domain.IAccountRepository.
func (account *Accounts) SaveInfo(
	ctx context.Context, acc domain.Account) error {
	if acc.Email == "" {
		return errors.New("missing key: email ")
	}
	if acc.Username != "" {
		if err := db.SafePgxWriteQuery(
			ctx, account.conn, UpdateAccountsUsername,
			acc.Username, acc.Email); err != nil {
			return err
		}

	}
	if acc.Password != "" {
		if err := db.SafePgxWriteQuery(
			ctx, account.conn, UpdateAccountsPassword,
			acc.Password, acc.Email); err != nil {
			return err
		}
	}
	if acc.Role != "" {
		if err := db.SafePgxWriteQuery(
			ctx, account.conn, UpdateAccountsRole,
			acc.Role, acc.Email); err != nil {
			return err
		}
	}
	return nil
}
