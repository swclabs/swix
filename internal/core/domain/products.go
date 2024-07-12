package domain

import "time"

// Products Table
type Products struct {
	ID          int64     `json:"id" db:"id"`
	Image       string    `json:"image" db:"image"`
	Price       string    `json:"price" db:"price"`
	Description string    `json:"description" db:"description"`
	Name        string    `json:"name" db:"name"`
	SupplierID  string    `json:"supplier_id" db:"supplier_id"`
	CategoryID  string    `json:"category_id" db:"category_id"`
	Spec        string    `json:"spec" db:"spec"`
	Status      string    `json:"status" db:"status"`
	Created     time.Time `json:"created" db:"created"`
}

// FavoriteProduct Table
type FavoriteProduct struct {
	ID        int64 `json:"id" db:"id"`
	UserID    int64 `json:"user_id" db:"user_id"`
	ProductID int64 `json:"product_id" db:"product_id"`
}

/*****************************************************************************/

type Specs struct {
	Screen  string `json:"screen"`
	Display string `json:"display"`
	SSD     []int  `json:"SSD"`
	RAM     []int  `json:"RAM"`
}

type Product struct {
	Specs
	Price       string `json:"price" validate:"required"`
	Description string `json:"description" validate:"required"`
	Name        string `json:"name" validate:"required"`
	SupplierID  string `json:"supplier_id" validate:"required"`
	CategoryID  string `json:"category_id" validate:"required"`
	Status      string `json:"status" validate:"required"`
}

type ProductSchema struct {
	ID          int64    `json:"id"`
	Image       []string `json:"image"`
	Price       string   `json:"price"`
	Description string   `json:"description"`
	Name        string   `json:"name"`
	Status      string   `json:"status"`
	Created     string   `json:"created"`
	Spec        Specs    `json:"spec"`
}

type UpdateProductInfo struct {
	Product
	ID int64 `json:"id" validate:"required"`
}

type CreateProductSchema struct {
	Msg string `json:"msg"`
	ID  int64  `json:"id"`
}
