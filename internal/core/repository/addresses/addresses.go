// Package repository
// Author: Duc Hung Ho @kieranhoo
package addresses

import (
	"context"
	"errors"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"

	"gorm.io/gorm"
)

type Addresses struct {
	conn *gorm.DB
}

func New(conn *gorm.DB) *Addresses {
	return &Addresses{
		conn: conn,
	}
}

// Insert implements domain.IAddressRepository.
func (addr *Addresses) Insert(ctx context.Context, data *domain.Addresses) error {
	if data == nil {
		return errors.New("input data invalid (nil)")
	}
	return db.SafeWriteQuery(
		ctx, addr.conn, InsertIntoAddresses,
		data.Street, data.Ward, data.District, data.City, data.Uuid,
	)
}
