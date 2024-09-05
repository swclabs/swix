package entity

import "time"

// Comment Table
type Comments struct {
	ID        int64     `json:"id" db:"id"`
	Level     int64     `json:"level" db:"level"`
	Content   string    `json:"content" db:"content"`
	UserID    int64     `json:"user_id" db:"user_id"`
	ProductID int64     `json:"product_id" db:"product_id"`
	ParentID  int64     `json:"parent_id" db:"parent_id"`
	Created   time.Time `json:"created" db:"created"`
}
