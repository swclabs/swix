package domain

// Carts table
type Carts struct {
	ID         int64 `json:"id" gorm:"column:id"`
	Quantity   int64 `json:"quantity" gorm:"column:quantity"`
	TotalPrice int64 `json:"total_price" gorm:"column:total_price"`
	UserID     int64 `json:"user_id" gorm:"column:user_id"`
}

// CartInfo schema
type CartInfo struct {
}

type ICartRepository interface {
	Add(productID int64) error
	AddMany(products []int64) error
	GetCartByUserID(userId int64) (*CartInfo, error)
	RemoveProduct(productID int64) error
}
