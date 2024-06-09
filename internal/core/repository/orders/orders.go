package orders

import (
	"gorm.io/gorm"
	"swclabs/swipecore/internal/core/domain"
)

type Orders struct {
	conn *gorm.DB
}

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

func New(conn *gorm.DB) *Orders {
	return &Orders{
		conn: conn,
	}
}

var _ IOrdersRepository = (*Orders)(nil)
