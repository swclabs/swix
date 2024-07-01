package orders

import (
	"context"
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

func (order *Orders) Create(ctx context.Context, userId string, cartId ...int64) error {
	//TODO implement me
	panic("implement me")
}
