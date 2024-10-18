package dtos

import "swclabs/swix/internal/core/domain/model"

// Cart schema request, response
type Cart struct {
	Name           string `json:"name"`
	CartID         int64  `json:"cart_id"`
	InventoryID    int64  `json:"inventory_id"`
	ProductID      int64  `json:"product_id"`
	Quantity       int64  `json:"quantity"`
	Color          string `json:"color"`
	InventoryPrice string `json:"price"`
	CurrencyCode   string `json:"currency_code"`
	InventoryImage string `json:"image"`
	InventorySpecs Specs  `json:"specs"`
	CategoryName   string `json:"category"`
}

// Carts schema
type Carts struct {
	UserID   int64  `json:"user_id"`
	Products []Cart `json:"products"`
}

// CartDTO request, response
type CartDTO struct {
	InventoryID int64 `json:"inventory_id" validate:"required"`
	Quantity    int64 `json:"quantity" validate:"required"`
}

type CartInsertDTO struct {
	CartDTO
	Email string `json:"email" validate:"required"`
}

// ProductOrderSchema is the schema for product in order
type ProductOrderSchema struct {
	ID             int64  `json:"id"`
	OrderID        int64  `json:"order_id"`
	ProductName    string `json:"product_name"`
	InventoryID    int64  `json:"inventory_id"`
	InventoryImage string `json:"inventory_image"`
	Color          string `json:"color"`
	Quantity       int64  `json:"quantity"`
	CurrencyCode   string `json:"currency_code"`
	TotalAmount    string `json:"total_amount"`
}

// OrderSchema is the schema response
type OrderSchema struct {
	ID        int64         `json:"id"`
	UUID      string        `json:"uuid"`
	Time      string        `json:"time"`
	Status    string        `json:"status"`
	UserID    int64         `json:"user_id"`
	Username  string        `json:"user_name"`
	UserEmail string        `json:"user_email"`
	Items     []model.Order `json:"items"`
}

// OrderDTO is the schema for creating an order request
type OrderDTO struct {
	DeleveryID int64 `json:"delevery_id" validate:"required"`
	Products   []struct {
		InventoryID int64 `json:"inventory_id" validate:"required"`
		Quantity    int64 `json:"quantity" validate:"required"`
		SpecsID     int64 `json:"specs_id"`
	} `json:"product" validate:"required"`
}

// CreateOrderSchema is the schema for creating an order request
type CreateOrderDTO struct {
	Email string `json:"email" validate:"required"`
	OrderDTO
}
