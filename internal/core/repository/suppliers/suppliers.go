package suppliers

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/infra/db"
)

// Suppliers struct for suppliers repository
type Suppliers struct {
	db db.IDatabase
}

// New creates a new Suppliers object.
func New(conn db.IDatabase) ISuppliersRepository {
	return useCache(&Suppliers{db: conn})
}

// Insert implements domain.ISuppliersRepository.
func (supplier *Suppliers) Insert(ctx context.Context, supp domain.Suppliers) error {
	return supplier.db.SafeWrite(
		ctx, insertIntoSuppliers, supp.Name, supp.Email)
}

// InsertAddress implements domain.ISuppliersRepository.
func (supplier *Suppliers) InsertAddress(
	ctx context.Context, addr domain.SuppliersAddress) error {
	return supplier.db.SafeWrite(
		ctx,
		insertIntoSuppliersAddress,
		addr.SuppliersID, addr.AddressUuiD,
	)
}

// GetLimit implements domain.ISuppliersRepository.
func (supplier *Suppliers) GetLimit(
	ctx context.Context, limit int) ([]domain.Suppliers, error) {
	// var _suppliers []domain.Suppliers
	rows, err := supplier.db.Query(ctx, selectSupplierByEmailLimit, limit)
	if err != nil {
		return nil, err
	}
	_suppliers, err := db.CollectRows[domain.Suppliers](rows)
	if err != nil {
		return nil, err
	}
	return _suppliers, nil
}

// GetByPhone implements domain.ISuppliersRepository.
func (supplier *Suppliers) GetByPhone(
	ctx context.Context, email string) (*domain.Suppliers, error) {
	// var _supplier domain.Suppliers
	rows, err := supplier.db.Query(ctx, selectByEmail, email)
	if err != nil {
		return nil, err
	}
	_supplier, err := db.CollectOneRow[domain.Suppliers](rows)
	if err != nil {
		return nil, err
	}
	return &_supplier, nil
}
