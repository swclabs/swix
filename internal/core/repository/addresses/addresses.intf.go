package addresses

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
)

// IAddress implements all methods to access the Addresses data in the database.
type IAddress interface {
	// Insert adds a new address to the database.
	// ctx is the context to manage the request's lifecycle.
	// data is a pointer to the Addresses object to be added.
	// Returns an error if any issues occur during the insertion process.
	Insert(ctx context.Context, data entity.Addresses) error

	GetByUserID(ctx context.Context, userID int64) ([]entity.Addresses, error)
}
