// Package addresses  Duc Hung Ho @kyeranyo
package addresses

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/pkg/infra/cache"
	"swclabs/swix/pkg/infra/db"
)

// New creates a new Addresses object
func New(conn db.IDatabase) IAddressRepository {
	return &Addresses{
		db: conn,
	}
}

// Addresses struct for address repository
type Addresses struct {
	db db.IDatabase
}

// GetByUserID implements IAddressRepository.
func (addr *Addresses) GetByUserID(ctx context.Context, userID int64) ([]entity.Addresses, error) {
	rows, err := addr.db.Query(ctx, selectAddressesByUserID, userID)
	if err != nil {
		return nil, err
	}
	addrData, err := db.CollectRows[entity.Addresses](rows)
	if err != nil {
		return nil, err
	}
	return addrData, nil
}

// Init initializes the Addresses object with database and redis connection
func Init(conn db.IDatabase, cache cache.ICache) IAddressRepository {
	return useCache(cache, New(conn))
}

// Insert implements IAddressRepository.
func (addr *Addresses) Insert(ctx context.Context, data entity.Addresses) error {
	return addr.db.SafeWrite(
		ctx, insertIntoAddresses,
		data.Street, data.Ward, data.District, data.City, data.UserID,
	)
}
