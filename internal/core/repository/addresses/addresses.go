// Package repository
// Author: Duc Hung Ho @kyeranyo
package addresses

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"

	"github.com/jackc/pgx/v5"
)

type Addresses struct {
	conn *pgx.Conn
}

func New(conn *pgx.Conn) *Addresses {
	return &Addresses{
		conn: conn,
	}
}

// Insert implements domain.IAddressRepository.
func (addr *Addresses) Insert(ctx context.Context, data domain.Addresses) error {
	return db.SafePgxWriteQuery(
		ctx, addr.conn, insertIntoAddresses,
		data.Street, data.Ward, data.District, data.City, data.Uuid,
	)
}
