package suppliers

import (
	"context"
	"swclabs/swix/app"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/pkg/infra/cache"
	"swclabs/swix/pkg/infra/db"
)

// New creates a new Suppliers object.
func New(conn db.IDatabase) ISuppliers {
	return &Suppliers{db: conn}
}

var _ = app.Repos(Init)

// Init initializes the Suppliers object with database and redis connection.
func Init(conn db.IDatabase, cache cache.ICache) ISuppliers {
	return useCache(cache, New(conn))
}

// Suppliers struct for suppliers repos
type Suppliers struct {
	db db.IDatabase
}

// Insert implements ISuppliersRepository.
func (supplier *Suppliers) Insert(ctx context.Context, supp entity.Suppliers) error {
	return supplier.db.SafeWrite(
		ctx, insertIntoSuppliers, supp.Name, supp.Email)
}

// GetLimit implements ISuppliersRepository.
func (supplier *Suppliers) GetLimit(
	ctx context.Context, limit int) ([]entity.Suppliers, error) {
	// var _suppliers []entity.Suppliers
	rows, err := supplier.db.Query(ctx, selectSupplierByEmailLimit, limit)
	if err != nil {
		return nil, err
	}
	_suppliers, err := db.CollectRows[entity.Suppliers](rows)
	if err != nil {
		return nil, err
	}
	return _suppliers, nil
}

// GetByPhone implements ISuppliersRepository.
func (supplier *Suppliers) GetByPhone(
	ctx context.Context, email string) (*entity.Suppliers, error) {
	// var _supplier entity.Suppliers
	rows, err := supplier.db.Query(ctx, selectByEmail, email)
	if err != nil {
		return nil, err
	}
	_supplier, err := db.CollectOneRow[entity.Suppliers](rows)
	if err != nil {
		return nil, err
	}
	return &_supplier, nil
}

// Edit implements ISuppliersRepository.
func (supplier *Suppliers) Edit(ctx context.Context, supp entity.Suppliers) error {
	return supplier.db.SafeWrite(
		ctx, updateSuppliers, supp.Name, supp.Email)
}
