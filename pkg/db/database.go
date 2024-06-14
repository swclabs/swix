package db

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5"

	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/pkg/utils"

	"go.uber.org/fx"
)

var (
	pgxConnection *pgx.Conn   = nil
	lock          *sync.Mutex = &sync.Mutex{}
	writeLock     *sync.Mutex = &sync.Mutex{}
)

func MigrateUp() error {
	const migrateUrl = "file://pkg/db/migration/"
	databaseUrl, err := utils.ConnectionURLBuilder("pg-migrate")
	if err != nil {
		return err
	}
	_migrate, err := migrate.New(migrateUrl, databaseUrl)
	if err != nil {
		return err
	}
	if err := _migrate.Up(); !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}

func MigrateDown() error {
	const migrateUrl = "file://pkg/db/migration/"
	databaseUrl, err := utils.ConnectionURLBuilder("pg-migrate")
	if err != nil {
		return err
	}
	_migrate, err := migrate.New(migrateUrl, databaseUrl)
	if err != nil {
		return err
	}
	if err := _migrate.Down(); !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}

func PgxConnection() *pgx.Conn {
	return pgxConnection
}

func CreateConnection(lc fx.Lifecycle, env config.Env) *pgx.Conn {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	var (
		dsn = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			env.DbUser, env.DbPassword, env.DbHost, env.DbPort, env.DbName)
		err error = nil
	)
	if pgxConnection == nil {
		lock.Lock()
		defer lock.Unlock()
		if pgxConnection == nil {
			pgxConnection, err = pgx.Connect(context.Background(), dsn)
		}
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err != nil {
				return err
			}
			fmt.Println("Connect to PostgreSQL successfully")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return pgxConnection.Close(ctx)
		},
	})
	return pgxConnection
}

func SafePgxWriteQuery(ctx context.Context, conn *pgx.Conn, sql string, args ...interface{}) error {
	// lock the connection
	writeLock.Lock()
	// after function call return, unlock the write lock
	defer writeLock.Unlock()
	_, err := conn.Exec(ctx, sql, args...)
	return err
}

func SafePgxWriteQueryReturnId(ctx context.Context, conn *pgx.Conn, sql string, args ...interface{}) (int64, error) {
	// lock the connection
	writeLock.Lock()
	// after function call return, unlock the write lock
	defer writeLock.Unlock()
	var id int64
	if err := conn.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}
