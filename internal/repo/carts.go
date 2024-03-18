// Package repo
// Author: Duc Hung Ho @kieranhoo
// Description: cart repository implementation
package repo

import (
	"github.com/swclabs/swipe-server/internal/domain"
	"log"

	"github.com/swclabs/swipe-server/pkg/db"
	"gorm.io/gorm"
)

type Carts struct {
	conn *gorm.DB
}

func NewCarts() domain.ICartRepository {
	_conn, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	return &Carts{
		conn: _conn,
	}
}

func (cart *Carts) Add(productID int64) error {
	//TODO implement me
	panic("implement me")
}

func (cart *Carts) AddMany(products []int64) error {
	//TODO implement me
	panic("implement me")
}

func (cart *Carts) GetCartByUserID(userId int64) (*domain.CartInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (cart *Carts) RemoveProduct(productID int64) error {
	//TODO implement me
	panic("implement me")
}
