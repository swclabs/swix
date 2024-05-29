// Package repository
// Author: Duc Hung Ho @kieranhoo
package addresses

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"
)

type Addresses struct {
	conn *gorm.DB
}

func NewAddresses() IAddressRepository {
	_conn, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	return &Addresses{
		conn: _conn,
	}
}

// Use implements domain.IAddressRepository.
func (addr *Addresses) Use(tx *gorm.DB) IAddressRepository {
	addr.conn = tx
	return addr
}

// Insert implements domain.IAddressRepository.
func (addr *Addresses) Insert(ctx context.Context, data *domain.Addresses) error {
	if data == nil {
		return errors.New("input data invalid (nil)")
	}
	return db.SafeWriteQuery(
		ctx,
		addr.conn,
		InsertIntoAddresses,
		data.Street, data.Ward, data.District, data.City, data.Uuid,
	)
}
