package suppliers

import (
	"context"
	"swclabs/swipex/internal/core/domain/entity"
)

// ISuppliers implements all methods to access and manage supplier data in the database.
type ISuppliers interface {
	// Insert adds a new supplier and their address to the database.
	// ctx is the context to manage the request's lifecycle.
	// sup is the Suppliers object to be added.
	// addr is the Addresses object associated with the supplier.
	// Returns an error if any issues occur during the insertion process.
	Insert(ctx context.Context, sup entity.Supplier) error

	// GetLimit retrieves a list of suppliers with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of suppliers to retrieve.
	// Returns a slice of Suppliers objects and an error if any issues occur during the retrieval process.
	GetLimit(ctx context.Context, limit int) ([]entity.Supplier, error)

	// GetByPhone retrieves a supplier by their phone number.
	// ctx is the context to manage the request's lifecycle.
	// phone is the phone number to search for.
	// Returns a pointer to the Suppliers object and an error if any issues occur during the retrieval process.
	GetByPhone(ctx context.Context, phone string) (*entity.Supplier, error)

	// Edit updates an existing supplier's information in the database.
	// ctx is the context to manage the request's lifecycle.
	// sup is the Suppliers object to be updated.
	// Returns an error if any issues occur during the update process.
	Edit(ctx context.Context, sup entity.Supplier) error
}
