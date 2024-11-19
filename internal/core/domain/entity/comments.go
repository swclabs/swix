package entity

import "time"

// Comment Table
type Comment struct {
	ID          int64     `json:"id" db:"id"`
	Content     string    `json:"content" db:"content"`
	UserID      int64     `json:"user_id" db:"user_id"`
	ProductID   int64     `json:"product_id" db:"product_id"`
	InventoryID int64     `json:"inventory_id" db:"inventory_id"`
	StarID      int64     `json:"star_id" db:"star_id"`
	Created     time.Time `json:"created" db:"created"`
}
