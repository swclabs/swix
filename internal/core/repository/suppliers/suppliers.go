package suppliers

import (
	"context"
	"log"
	"swclabs/swipecore/internal/core/repository/addresses"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"
)

type Suppliers struct {
	conn *gorm.DB
}

func NewSuppliers() ISuppliersRepository {
	_conn, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	return &Suppliers{
		conn: _conn,
	}
}

// Use implements domain.ISuppliersRepository.
func (supplier *Suppliers) Use(tx *gorm.DB) ISuppliersRepository {
	supplier.conn = tx
	return supplier
}

// Insert implements domain.ISuppliersRepository.
func (supplier *Suppliers) Insert(ctx context.Context, supp domain.Suppliers, addr domain.Addresses) error {
	return supplier.conn.Transaction(func(tx *gorm.DB) error {
		if err := db.SafeWriteQuery(
			ctx,
			tx,
			InsertIntoSuppliers,
			supp.Name, supp.Email,
		); err != nil {
			return err
		}
		_supplier, err := NewSuppliers().Use(tx).GetByPhone(ctx, supp.Email)
		if err != nil {
			return err
		}
		addr.Uuid = uuid.New().String()
		if err := addresses.NewAddresses().Use(tx).Insert(ctx, &addr); err != nil {
			return err
		}
		return NewSuppliers().Use(tx).InsertAddress(ctx, domain.SuppliersAddress{
			SuppliersID: _supplier.Id,
			AddressUuiD: addr.Uuid,
		})
	})
}

// InsertAddress implements domain.ISuppliersRepository.
func (supplier *Suppliers) InsertAddress(ctx context.Context, addr domain.SuppliersAddress) error {
	return db.SafeWriteQuery(
		ctx,
		supplier.conn,
		InsertIntoSuppliersAddress,
		addr.SuppliersID, addr.AddressUuiD,
	)
}

// GetLimit implements domain.ISuppliersRepository.
func (supplier *Suppliers) GetLimit(ctx context.Context, limit int) ([]domain.Suppliers, error) {
	var _suppliers []domain.Suppliers
	if err := supplier.conn.
		Table(domain.SuppliersTable).Find(&_suppliers).Limit(limit).Error; err != nil {
		return nil, err
	}
	return _suppliers, nil
}

// GetByPhone implements domain.ISuppliersRepository.
func (supplier *Suppliers) GetByPhone(ctx context.Context, email string) (*domain.Suppliers, error) {
	var _supplier domain.Suppliers
	if err := supplier.conn.
		Table(domain.SuppliersTable).Where("email = ?", email).First(&_supplier).Error; err != nil {
		return nil, err
	}
	return &_supplier, nil
}
