// Package addresses  Duc Hung Ho @kyeranyo
package addresses

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/pkg/infra/cache"
	"swclabs/swipecore/pkg/infra/db"
)

// Addresses struct for address repository
type Addresses struct {
	db db.IDatabase
}

// New creates a new Addresses object
func New(conn db.IDatabase) IAddressRepository {
	return &Addresses{
		db: conn,
	}
}

// Init initializes the Addresses object with database and redis connection
func Init(conn db.IDatabase, cache cache.ICache) IAddressRepository {
	return useCache(cache, New(conn))
}

// Insert implements IAddressRepository.
func (addr *Addresses) Insert(ctx context.Context, data entity.Addresses) error {
	return addr.db.SafeWrite(
		ctx, insertIntoAddresses,
		data.Street, data.Ward, data.District, data.City, data.UUID,
	)
}
