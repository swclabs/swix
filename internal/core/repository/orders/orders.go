package orders

import (
	"swclabs/swipecore/internal/core/domain"

	"github.com/jackc/pgx/v5"
)

type Orders struct {
	conn *pgx.Conn
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

func New(conn *pgx.Conn) *Orders {
	return &Orders{
		conn: conn,
	}
}

var _ IOrdersRepository = (*Orders)(nil)
