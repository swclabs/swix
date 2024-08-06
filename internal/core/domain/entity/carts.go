// Package entity Cart entities
package entity

// Carts table
type Carts struct {
	ID          int64 `json:"id" db:"id"`
	UserID      int64 `json:"user_id" db:"user_id"`
	InventoryID int64 `json:"inventory_id" db:"inventory_id"`
	Quantity    int64 `json:"quantity" db:"quantity"`
	SpecID      int64 `json:"spec_id" db:"spec_id"`
}
