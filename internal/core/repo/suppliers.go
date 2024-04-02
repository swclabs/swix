package repo

import (
	"context"
	"log"

	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/pkg/db"
	"swclabs/swipe-api/pkg/db/queries"

	"gorm.io/gorm"
)

type Suppliers struct {
	data *domain.Suppliers
	conn *gorm.DB
}

func NewSuppliers() domain.ISuppliersRepository {
	_conn, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	return &Suppliers{
		data: &domain.Suppliers{},
		conn: _conn,
	}
}

func (supplier *Suppliers) New(ctx context.Context, supp *domain.Suppliers, addr *domain.Addresses) error {
	return supplier.conn.Transaction(func(tx *gorm.DB) error {
		return db.SafeWriteQuery(
			ctx,
			supplier.conn,
			queries.InsertIntoSuppliers,
			supp.Name, supp.PhoneNumber, supp.Email,
		)
	})
}

func (supplier *Suppliers) GetAll(ctx context.Context) ([]domain.Suppliers, error) {
	panic("not implemented")
}
