package repo

import (
	"log"

	"github.com/swclabs/swipe-api/internal/core/domain"
	"github.com/swclabs/swipe-api/pkg/db"
	"github.com/swclabs/swipe-api/pkg/db/queries"
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

func (supplier *Suppliers) New(supp *domain.Suppliers, addr *domain.Addresses) error {
	return supplier.conn.Transaction(func(tx *gorm.DB) error {
		return db.SafeWriteQuery(supplier.conn,
			queries.InsertIntoSuppliers,
			supp.Name, supp.PhoneNumber, supp.Email,
		)
	})
}