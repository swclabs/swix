// Package db connect to database
package db

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5"
)

// Transaction is a wrapper for pgx.Tx
type Transaction struct {
	tx        pgx.Tx
	lock      *sync.Mutex
	writeLock *sync.Mutex
}

var _ ITransaction = (*Transaction)(nil)

// NewTx implements IDatabase.
func NewTx(ctx context.Context) (ITransaction, error) {
	tx, err := pgxConnection.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return &Transaction{
		tx:        tx,
		lock:      lock,
		writeLock: writeLock,
	}, nil
}

// Query implements ITransaction.
func (t *Transaction) Query(ctx context.Context, sql string, args ...interface{}) (Rows, error) {
	return t.tx.Query(ctx, sql, args...)
}

// Rollback implements ITransaction.
func (t *Transaction) Rollback(ctx context.Context) error {
	return t.tx.Rollback(ctx)
}

// Commit implements ITransaction.
func (t *Transaction) Commit(ctx context.Context) error {
	if err := t.tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

// SafeWrite implements ITransaction.
func (t *Transaction) SafeWrite(ctx context.Context, sql string, args ...interface{}) error {
	t.writeLock.Lock()
	// after function call return, unlock the write lock
	defer t.writeLock.Unlock()
	_, err := t.tx.Conn().Exec(ctx, sql, args...)
	return err
}

// SafeWriteReturn implements ITransaction.
func (t *Transaction) SafeWriteReturn(ctx context.Context, sql string, args ...interface{}) (int64, error) {
	// lock the connection
	t.writeLock.Lock()
	// after function call return, unlock the write lock
	defer t.writeLock.Unlock()
	var id int64
	if err := t.tx.Conn().QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}
