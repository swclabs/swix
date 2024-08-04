package entity

// Specifications Table
type Specifications struct {
	ID          int64  `json:"id" db:"id"`
	InventoryID int64  `json:"inventory_id" db:"inventory_id"`
	Content     string `json:"content" db:"content"`
}
