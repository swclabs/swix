package db

import (
	"context"
	"errors"
	"sync"

	"github.com/golang-migrate/migrate/v4"

	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/pkg/utils"

	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB = nil
var lock *sync.Mutex = &sync.Mutex{}
var writeLock *sync.Mutex = &sync.Mutex{}

func Connection(env config.Env) (*gorm.DB, error) {
	dsn, _ := utils.ConnectionURLBuilderWithEnv("postgres", env)
	if connection == nil {
		lock.Lock()
		defer lock.Unlock()
		if connection == nil {
			conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err != nil {
				return nil, err
			}
			connection = conn
		}
	}
	return connection, nil
	// return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func DatabaseConnection() *gorm.DB {
	return connection
}

func CreatePostgresConnection(lc fx.Lifecycle, env config.Env) *gorm.DB {
	dsn, _ := utils.ConnectionURLBuilderWithEnv("postgres", env)
	var err error = nil
	if connection == nil {
		lock.Lock()
		defer lock.Unlock()
		if connection == nil {
			connection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		}
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err != nil {
				return err
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
	return connection
}

func SafeWriteQuery(ctx context.Context, connection *gorm.DB, sql string, args ...interface{}) error {
	// lock the connection
	writeLock.Lock()
	// after function call return, unlock the write lock
	defer writeLock.Unlock()
	return connection.WithContext(ctx).Exec(sql, args...).Error
}

func SafeWriteQueryReturnId(ctx context.Context, connection *gorm.DB, sql string, args ...interface{}) (int64, error) {
	// lock the connection
	writeLock.Lock()
	// after function call return, unlock the write lock
	defer writeLock.Unlock()
	var id int64
	if err := connection.WithContext(ctx).Raw(sql, args...).Scan(&id).Error; err != nil {
		return -1, err
	}
	return id, nil
}

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
