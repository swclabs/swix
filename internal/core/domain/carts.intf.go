package domain

import "context"

type ICartRepository interface {
	Insert(ctx context.Context, productID int64) error
	InsertMany(ctx context.Context, products []int64) error
	GetCartByUserID(ctx context.Context, userId int64) (*CartInfo, error)
	RemoveProduct(ctx context.Context, productID int64, userId int64) error
}
