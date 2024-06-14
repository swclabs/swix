package addresses

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

// IAddressRepository implements all methods to access the Addresses data in the database.
type IAddressRepository interface {
	// Insert adds a new address to the database.
	// ctx is the context to manage the request's lifecycle.
	// data is a pointer to the Addresses object to be added.
	// Returns an error if any issues occur during the insertion process.
	Insert(ctx context.Context, data domain.Addresses) error
}
