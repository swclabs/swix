// Package repo
// Author: Duc Hung Ho @kieranhoo
// Description: cart repository implementation
package repo

import (
	"log"

	"github.com/swclabs/swipe-api/pkg/db"
	"gorm.io/gorm"
)

type Carts struct {
	conn *gorm.DB
}

func NewCarts() *Carts {
	_conn, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	return &Carts{
		conn: _conn,
	}
}
