// Package entity Cart entities
package entity

// Cart table
type Cart struct {
	ID          int64 `json:"id" db:"id"`
	UserID      int64 `json:"user_id" db:"user_id"`
	InventoryID int64 `json:"inventory_id" db:"inventory_id"`
	Quantity    int64 `json:"quantity" db:"quantity"`
}
