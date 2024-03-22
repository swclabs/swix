// Package repo
// Author: Duc Hung Ho @kieranhoo
package repo

import (
	"errors"
	"log"

	"github.com/swclabs/swipe-api/internal/core/domain"
	"github.com/swclabs/swipe-api/pkg/db"
	"github.com/swclabs/swipe-api/pkg/db/queries"
	"gorm.io/gorm"
)

type Addresses struct {
	data *domain.Addresses
	conn *gorm.DB
}

func NewAddresses() domain.IAddressRepository {
	_conn, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}
	return &Addresses{
		data: &domain.Addresses{},
		conn: _conn,
	}
}

func (addr *Addresses) New(data *domain.Addresses) error {
	if data == nil {
		return errors.New("input data invalid (nil)")
	}
	if (data.SupplierId != "" && data.UserId != "") ||
		(data.SupplierId == "" && data.UserId == "") {
		return errors.New("unknown data from supplier or user")
	}
	return db.SafeWriteQuery(addr.conn,
		queries.InsertIntoAddresses,
		data.UserId, data.SupplierId, data.Street, data.Ward, data.District, data.City,
	)
}
