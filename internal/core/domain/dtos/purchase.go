package dtos

import (
	"github.com/swclabs/swipex/internal/core/domain/model"
)

// Cart schema request, response
type Cart struct {
	Name           string `json:"name"`
	CartID         int64  `json:"cart_id"`
	InventoryID    int64  `json:"inventory_id"`
	ProductID      int64  `json:"product_id"`
	Quantity       int64  `json:"quantity"`
	Color          string `json:"color"`
	Code           string `json:"code"`
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

type OrderInfo struct {
	UUID          string            `json:"uuid"`
	CreatedAt     string            `json:"time"`
	PaymentMethod string            `json:"payment_method"`
	User          OrderFormCustomer `json:"user"`
	Delivery      OrderFormDelivery `json:"delivery"`
	Address       OrderFormAddress  `json:"address"`
	TotalAmount   string            `json:"total_amount"`
	Items         []model.Order     `json:"items"`
}

type Order struct {
	CouponCode    string             `json:"coupon_code"`
	PaymentMethod string             `json:"payment_method" validate:"required"`
	Customer      OrderFormCustomer  `json:"customer" validate:"required"`
	Delivery      OrderFormDelivery  `json:"delivery" validate:"required"`
	Address       OrderFormAddress   `json:"address" validate:"required"`
	Product       []OrderFormProduct `json:"product" validate:"required"`
}

type OrderFormAddress struct {
	City     string `json:"city" validate:"required"`
	Ward     string `json:"ward" validate:"required"`
	District string `json:"district" validate:"required"`
	Street   string `json:"street" validate:"required"`
}

type OrderFormProduct struct {
	Code     string `json:"code" validate:"required"`
	Quantity int64  `json:"quantity" validate:"required"`
}

type OrderFormDelivery struct {
	Status   string `json:"status" validate:"required"`
	Method   string `json:"method" validate:"required"`
	Note     string `json:"note" `
	SentDate string `json:"sent_date"`
}

type OrderFormCustomer struct {
	Email     string `json:"email" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Phone     string `json:"phone" validate:"required,number"`
}
type OrderForm struct {
	CouponCode    string             `json:"coupon_code"`
	PaymentMethod string             `json:"payment_method" validate:"required"`
	Customer      OrderFormCustomer  `json:"customer" validate:"required"`
	Delivery      OrderFormDelivery  `json:"delivery" validate:"required"`
	Address       OrderFormAddress   `json:"address" validate:"required"`
	Product       []OrderFormProduct `json:"product" validate:"required"`
}

type OrderStatus struct {
	OrderCode string `json:"order_code" validate:"required"`
	Status    string `json:"status" validate:"required"`
}

type OrderResponse struct {
	OrderCode string `json:"order_code"`
}
