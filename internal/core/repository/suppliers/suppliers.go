package suppliers

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/repository/addresses"
	"swclabs/swipecore/pkg/db"

	"github.com/google/uuid"
)

type Suppliers struct {
	db db.IDatabase
}

func New(conn db.IDatabase) ISuppliersRepository {
	return useCache(&Suppliers{db: conn})
}

// Insert implements domain.ISuppliersRepository.
func (supplier *Suppliers) Insert(
	ctx context.Context, supp domain.Suppliers, addr domain.Addresses) error {

	tx, err := db.BeginTransaction(ctx)
	if err != nil {
		return err
	}

	var errTx error = nil
	defer func() {
		if errTx != nil {
			// Sentry Capture failed
			tx.Rollback(ctx)
		}
	}()

	if errTx = tx.SafeWrite(
		ctx, insertIntoSuppliers, supp.Name, supp.Email); errTx != nil {
		return errTx
	}

	_supplier, errTx := New(tx).GetByPhone(ctx, supp.Email)
	if errTx != nil {
		return errTx
	}

	addr.Uuid = uuid.New().String()
	if errTx = addresses.New(tx).Insert(ctx, addr); errTx != nil {
		return errTx
	}

	errTx = New(tx).InsertAddress(ctx, domain.SuppliersAddress{
		SuppliersID: _supplier.Id,
		AddressUuiD: addr.Uuid,
	})

	return tx.Commit(ctx)
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
