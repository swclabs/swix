// Package addresses  Duc Hung Ho @kyeranyo
package addresses

import (
	"context"

	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/pkg/infra/cache"
	"github.com/swclabs/swipex/pkg/infra/db"
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
func (addr *Addresses) GetByID(ctx context.Context, id int64) (*entity.Address, error) {
	row, err := addr.db.Query(ctx, selectAddressesByID, id)
	if err != nil {
		return nil, err
	}
	addrData, err := db.CollectRow[entity.Address](row)
	return &addrData, nil
}

// GetByUserID implements IAddressRepository.
func (addr *Addresses) GetByUserID(ctx context.Context, userID int64) ([]entity.Address, error) {
	rows, err := addr.db.Query(ctx, selectAddressesByUserID, userID)
	if err != nil {
		return nil, err
	}
	addrData, err := db.CollectRows[entity.Address](rows)
	if err != nil {
		return nil, err
	}
	return addrData, nil
}

// Insert implements IAddressRepository.
func (addr *Addresses) Insert(ctx context.Context, data entity.Address) (int64, error) {
	return addr.db.SafeWriteReturn(
		ctx, insertIntoAddresses,
		data.Street, data.Ward, data.District, data.City, data.UserID,
	)
}
