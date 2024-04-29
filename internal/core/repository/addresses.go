// Package repository
// Author: Duc Hung Ho @kieranhoo
package repository

import (
	"context"
	"errors"
	"log"

	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"
	"swclabs/swipecore/pkg/db/queries"

	"gorm.io/gorm"
)

type Addresses struct {
	conn *gorm.DB
}

func NewAddresses() domain.IAddressRepository {
	_conn, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	return &Addresses{
		conn: _conn,
	}
}

func (addr *Addresses) Use(tx *gorm.DB) domain.IAddressRepository {
	addr.conn = tx
	return addr
}

func (addr *Addresses) Insert(ctx context.Context, data *domain.Addresses) error {
	if data == nil {
		return errors.New("input data invalid (nil)")
	}
	return db.SafeWriteQuery(
		ctx,
		addr.conn,
		queries.InsertIntoAddresses,
		data.Street, data.Ward, data.District, data.City, data.Uuid,
	)
}
