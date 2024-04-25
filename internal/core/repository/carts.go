// Package repository
// Author: Duc Hung Ho @kieranhoo
// Description: cart repository implementation
package repository

import (
	"log"
	"swclabs/swipe-api/internal/core/domain"

	"gorm.io/gorm"
	"swclabs/swipe-api/pkg/db"
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

func (cart *Carts) Insert(productID int64) error {
	//TODO implement me
	panic("implement me")
}

func (cart *Carts) InsertMany(products []int64) error {
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