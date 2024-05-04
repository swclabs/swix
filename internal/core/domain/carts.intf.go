package domain

type ICartRepository interface {
	Insert(productID int64) error
	InsertMany(products []int64) error
	GetCartByUserID(userId int64) (*CartInfo, error)
	RemoveProduct(productID int64) error
}
