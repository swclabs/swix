package repository

import (
	"context"
	"log"

	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"
	"swclabs/swipecore/pkg/db/queries"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Suppliers struct {
	conn *gorm.DB
}

func NewSuppliers() domain.ISuppliersRepository {
	_conn, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	return &Suppliers{
		conn: _conn,
	}
}

func (supplier *Suppliers) Use(tx *gorm.DB) domain.ISuppliersRepository {
	supplier.conn = tx
	return supplier
}

func (supplier *Suppliers) Insert(ctx context.Context, supp domain.Suppliers, addr domain.Addresses) error {
	return supplier.conn.Transaction(func(tx *gorm.DB) error {
		if err := db.SafeWriteQuery(
			ctx,
			tx,
			queries.InsertIntoSuppliers,
			supp.Name, supp.PhoneNumber, supp.Email,
		); err != nil {
			return err
		}
		_supplier, err := NewSuppliers().Use(tx).GetByPhone(ctx, supp.Email)
		if err != nil {
			return err
		}
		addr.Uuid = uuid.New().String()
		if err := NewAddresses().Use(tx).Insert(ctx, &addr); err != nil {
			return err
		}
		return NewSuppliers().Use(tx).InsertAddress(ctx, domain.SuppliersAddress{
			SuppliersID: _supplier.Id,
			AddressUuiD: addr.Uuid,
		})
	})
}

func (supplier *Suppliers) InsertAddress(ctx context.Context, addr domain.SuppliersAddress) error {
	return db.SafeWriteQuery(
		ctx,
		supplier.conn,
		queries.InsertIntoSuppliersAddress,
		addr.SuppliersID, addr.AddressUuiD,
	)
}

func (supplier *Suppliers) GetLimit(ctx context.Context, limit int) ([]domain.Suppliers, error) {
	var _suppliers []domain.Suppliers
	if err := supplier.conn.Table(domain.SuppliersTable).Find(&_suppliers).Limit(limit).Error; err != nil {
		return nil, err
	}
	return _suppliers, nil
}

func (supplier *Suppliers) GetByPhone(ctx context.Context, email string) (*domain.Suppliers, error) {
	var _supplier domain.Suppliers
	if err := supplier.conn.Table(domain.SuppliersTable).Where("email = ?", email).First(&_supplier).Error; err != nil {
		return nil, err
	}
	return &_supplier, nil
}
