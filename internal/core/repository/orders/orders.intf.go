package orders

import "swclabs/swipecore/internal/core/domain"

type IOrdersRepository interface {
	Create() error
	GetById(id string) (domain.Orders, error)
	GetAll() ([]domain.Orders, error)
	Delete(id string) error
}
