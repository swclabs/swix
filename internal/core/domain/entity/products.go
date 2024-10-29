package entity

import "time"

// Product Table
type Product struct {
	ID          int64     `json:"id" db:"id"`
	Image       string    `json:"image" db:"image"`
	ShopImage   string    `json:"shop_image" db:"shop_image"`
	Price       string    `json:"price" db:"price"`
	Description string    `json:"description" db:"description"`
	Name        string    `json:"name" db:"name"`
	SupplierID  int64     `json:"supplier_id" db:"supplier_id"`
	CategoryID  int64     `json:"category_id" db:"category_id"`
	Specs       string    `json:"specs" db:"specs"`
	Status      string    `json:"status" db:"status"`
	Created     time.Time `json:"created" db:"created"`
	Rating      float64   `json:"rating" db:"rating"`
}
