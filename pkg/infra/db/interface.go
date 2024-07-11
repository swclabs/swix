package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type IDatabase interface {
	Query(ctx context.Context, sql string, args ...interface{}) (Rows, error)
	SafeWrite(ctx context.Context, sql string, args ...interface{}) error
	SafeWriteReturn(ctx context.Context, sql string, args ...interface{}) (int64, error)
}

type ITransaction interface {
	IDatabase
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type Rows pgx.Rows
