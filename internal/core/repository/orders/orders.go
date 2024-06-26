package orders

import (
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"
)

type Orders struct {
	db db.IDatabase
}

func New(conn db.IDatabase) IOrdersRepository {
	return &Orders{
		db: conn,
	}
}

var _ IOrdersRepository = (*Orders)(nil)

func (order *Orders) Create() error {
	//TODO implement me
	panic("implement me")
}

func (order *Orders) GetById(id string) (domain.Orders, error) {
	//TODO implement me
	panic("implement me")
}

func (order *Orders) GetAll() ([]domain.Orders, error) {
	//TODO implement me
	panic("implement me")
}

func (order *Orders) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
