package model

// Comment Table comments
type Comment struct {
	ID        int64  `json:"id" gorm:"column:id"`
	Level     int64  `json:"level" gorm:"column:level"`
	Content   string `json:"content" gorm:"column:content"`
	UserId    string `json:"user_id" gorm:"column:user_id"`
	ProductID string `json:"product_id" gorm:"column:product_id"`
}
