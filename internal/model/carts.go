package model

// Cart table Carts
type Cart struct {
	ID         int64 `json:"id" gorm:"column:id"`
	Quantity   int64 `json:"quantity" gorm:"column:quantity"`
	TotalPrice int64 `json:"total_price" gorm:"column:total_price"`
	UserID     int64 `json:"user_id" gorm:"column:user_id"`
}
