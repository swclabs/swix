package domain

import (
	"context"

	"gorm.io/gorm"
)

// IAddressRepository implements all methods to access the Addresses data in the database.
type IAddressRepository interface {
	// Use sets the transaction connection.
	// tx is the transaction connection using gorm.DB.
	// Returns an instance of IAddressRepository with the transaction set.
	Use(tx *gorm.DB) IAddressRepository

	// Insert adds a new address to the database.
	// ctx is the context to manage the request's lifecycle.
	// data is a pointer to the Addresses object to be added.
	// Returns an error if any issues occur during the insertion process.
	Insert(ctx context.Context, data *Addresses) error
}
