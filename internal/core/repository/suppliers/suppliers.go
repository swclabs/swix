package suppliers

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/repository/addresses"
	"swclabs/swipecore/pkg/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Suppliers struct {
	conn *pgx.Conn
}

func New(conn *pgx.Conn) *Suppliers {
	return &Suppliers{conn: conn}
}

// Insert implements domain.ISuppliersRepository.
func (supplier *Suppliers) Insert(
	ctx context.Context, supp domain.Suppliers, addr domain.Addresses) error {

	tx, err := supplier.conn.Begin(ctx)
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

	if errTx = db.SafePgxWriteQuery(
		ctx, tx.Conn(), insertIntoSuppliers, supp.Name, supp.Email); errTx != nil {
		return errTx
	}

	_supplier, errTx := New(tx.Conn()).GetByPhone(ctx, supp.Email)
	if errTx != nil {
		return errTx
	}

	addr.Uuid = uuid.New().String()
	if errTx = addresses.New(tx.Conn()).Insert(ctx, addr); errTx != nil {
		return errTx
	}

	errTx = New(tx.Conn()).InsertAddress(ctx, domain.SuppliersAddress{
		SuppliersID: _supplier.Id,
		AddressUuiD: addr.Uuid,
	})

	return tx.Commit(ctx)
}

// InsertAddress implements domain.ISuppliersRepository.
func (supplier *Suppliers) InsertAddress(
	ctx context.Context, addr domain.SuppliersAddress) error {
	return db.SafePgxWriteQuery(
		ctx, supplier.conn,
		insertIntoSuppliersAddress,
		addr.SuppliersID, addr.AddressUuiD,
	)
}

// GetLimit implements domain.ISuppliersRepository.
func (supplier *Suppliers) GetLimit(
	ctx context.Context, limit int) ([]domain.Suppliers, error) {
	// var _suppliers []domain.Suppliers
	rows, err := supplier.conn.Query(ctx, selectSupplierByEmailLimit, limit)
	if err != nil {
		return nil, err
	}
	_suppliers, err := pgx.CollectRows[domain.Suppliers](rows, pgx.RowToStructByName[domain.Suppliers])
	if err != nil {
		return nil, err
	}
	return _suppliers, nil
}

// GetByPhone implements domain.ISuppliersRepository.
func (supplier *Suppliers) GetByPhone(
	ctx context.Context, email string) (*domain.Suppliers, error) {
	// var _supplier domain.Suppliers
	rows, err := supplier.conn.Query(ctx, selectByEmail, email)
	if err != nil {
		return nil, err
	}
	_supplier, err := pgx.CollectOneRow[domain.Suppliers](rows, pgx.RowToStructByName[domain.Suppliers])
	if err != nil {
		return nil, err
	}
	return &_supplier, nil
}
