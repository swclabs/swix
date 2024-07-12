package domain

// Comment Table
type Comment struct {
	ID        int64  `json:"id" db:"id"`
	Level     int64  `json:"level" db:"level"`
	Content   string `json:"content" db:"content"`
	UserID    string `json:"user_id" db:"user_id"`
	ProductID string `json:"product_id" db:"product_id"`
}
