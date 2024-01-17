package db

import (
	"sync"

	"github.com/swclabs/swipe-api/pkg/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB = nil
var lock *sync.Mutex = &sync.Mutex{}

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
