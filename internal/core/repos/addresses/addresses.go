// Package addresses  Duc Hung Ho @kyeranyo
package addresses

import (
	"context"
	"swclabs/swix/app"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/pkg/infra/cache"
	"swclabs/swix/pkg/infra/db"
)

var _ = app.Repos(Init)

// New creates a new Addresses object
func New(conn db.IDatabase) IAddress {
	return &Addresses{
		db: conn,
	}
}

// Init initializes the Addresses object with database and redis connection
func Init(conn db.IDatabase, cache cache.ICache) IAddress {
	return useCache(cache, New(conn))
}

// Addresses struct for address repos
type Addresses struct {
	db db.IDatabase
}

// GetByID implements IAddress.
func (addr *Addresses) GetByID(ctx context.Context, id int64) (*entity.Addresses, error) {
	row, err := addr.db.Query(ctx, selectAddressesByID, id)
	if err != nil {
		return nil, err
	}
	addrData, err := db.CollectRow[entity.Addresses](row)
	return &addrData, nil
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

// Insert implements IAddressRepository.
func (addr *Addresses) Insert(ctx context.Context, data entity.Addresses) (int64, error) {
	return addr.db.SafeWriteReturn(
		ctx, insertIntoAddresses,
		data.Street, data.Ward, data.District, data.City, data.UserID,
	)
}
