package orders

import "context"

type IOrdersRepository interface {
	Create(ctx context.Context, userId string, cartId ...int64) error
}
