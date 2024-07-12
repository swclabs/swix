// Package db connect to database
package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// IDatabase interface for database
type IDatabase interface {
	Query(ctx context.Context, sql string, args ...interface{}) (Rows, error)
	SafeWrite(ctx context.Context, sql string, args ...interface{}) error
	SafeWriteReturn(ctx context.Context, sql string, args ...interface{}) (int64, error)
}

// ITransaction interface for transaction
type ITransaction interface {
	IDatabase
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

// Rows is a type alias for pgx.Rows
type Rows pgx.Rows
