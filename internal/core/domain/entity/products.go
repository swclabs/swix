package entity

import "time"

// Products Table
type Products struct {
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
}

// FavoriteProduct Table
type FavoriteProduct struct {
	ID        int64 `json:"id" db:"id"`
	UserID    int64 `json:"user_id" db:"user_id"`
	ProductID int64 `json:"product_id" db:"product_id"`
}
