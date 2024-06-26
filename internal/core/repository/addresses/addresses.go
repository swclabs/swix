// Package repository
// Author: Duc Hung Ho @kyeranyo
package addresses

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"
)

type Addresses struct {
	db db.IDatabase
}

func New(conn db.IDatabase) IAddressRepository {
	return &Addresses{
		db: conn,
	}
}

// Insert implements domain.IAddressRepository.
func (addr *Addresses) Insert(ctx context.Context, data domain.Addresses) error {
	return addr.db.SafeWrite(
		ctx, insertIntoAddresses,
		data.Street, data.Ward, data.District, data.City, data.Uuid,
	)
}
