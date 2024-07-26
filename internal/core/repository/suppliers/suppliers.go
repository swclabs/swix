package suppliers

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/pkg/infra/cache"
	"swclabs/swipecore/pkg/infra/db"
)

// Suppliers struct for suppliers repository
type Suppliers struct {
	db db.IDatabase
}

// New creates a new Suppliers object.
func New(conn db.IDatabase) ISuppliersRepository {
	return &Suppliers{db: conn}
}

// Init initializes the Suppliers object with database and redis connection.
func Init(conn db.IDatabase, cache cache.ICache) ISuppliersRepository {
	return useCache(cache, New(conn))
}

// Insert implements ISuppliersRepository.
func (supplier *Suppliers) Insert(ctx context.Context, supp entity.Suppliers) error {
	return supplier.db.SafeWrite(
		ctx, insertIntoSuppliers, supp.Name, supp.Email)
}

// InsertAddress implements ISuppliersRepository.
func (supplier *Suppliers) InsertAddress(
	ctx context.Context, addr entity.SuppliersAddress) error {
	return supplier.db.SafeWrite(
		ctx,
		insertIntoSuppliersAddress,
		addr.SuppliersID, addr.AddressUuiD,
	)
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

// Edit implements domain.ISuppliersRepository.
func (supplier *Suppliers) Edit(ctx context.Context, supp domain.Suppliers) error {
	return supplier.db.SafeWrite(
		ctx, updateSuppliers, supp.Name, supp.Email)
}
