package db

import (
	"context"
	"sync"

	"swclabs/swipecore/pkg/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB = nil
var lock *sync.Mutex = &sync.Mutex{}
var writeLock *sync.Mutex = &sync.Mutex{}

func Connection() (*gorm.DB, error) {
	dsn, _ := utils.ConnectionURLBuilder("postgres")
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

func SafeWriteQuery(ctx context.Context, connection *gorm.DB, sql string, args ...interface{}) error {
	// lock the connection
	writeLock.Lock()
	// after function call return, unlock the write lock
	defer writeLock.Unlock()
	return connection.WithContext(ctx).Exec(sql, args...).Error
}
