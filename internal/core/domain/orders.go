package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

/*
CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "uuid" varchar NOT NULL,
  "time" timestamptz NOT NULL,
  "user_id" bigint NOT NULL,
  "in_word" bigint NOT NULL,
  "status" varchar NOT NULL
);

CREATE TABLE "product_in_order" (
  "id" bigserial PRIMARY KEY,
  "order_id" bigint NOT NULL,
  "inventory_id" bigint NOT NULL,
  "quantity" bigint NOT NULL,
  "total_amount" NUMERIC(19, 4) NOT NULL,
  "currency_code" varchar(3) NOT NULL
);
*/

// Orders table schema
type Orders struct {
	ID          int64           `json:"id" db:"id"`
	UUID        string          `json:"uuid" db:"uuid"`
	UserID      int64           `json:"user_id" db:"user_id"`
	Status      string          `json:"status" db:"status"`
	TotalAmount decimal.Decimal `json:"total_amount" db:"total_amount"`
	Time        time.Time       `json:"time" db:"time"`
}

// ProductInOrder table schema
type ProductInOrder struct {
	ID           int64           `json:"id" db:"id"`
	OrderID      int64           `json:"order_id" db:"order_id"`
	InventoryID  int64           `json:"inventory_id" db:"inventory_id"`
	Quantity     int64           `json:"quantity" db:"quantity"`
	CurrencyCode string          `json:"currency_code" db:"currency_code"`
	TotalAmount  decimal.Decimal `json:"total_amount" db:"total_amount"`
}

// ProductOrderSchema is the schema for product in order
type ProductOrderSchema struct {
	ID           int64  `json:"id"`
	OrderID      int64  `json:"order_id"`
	InventoryID  int64  `json:"inventory_id" db:"inventory_id"`
	Quantity     int64  `json:"quantity"`
	CurrencyCode string `json:"currency_code"`
	TotalAmount  string `json:"total_amount"`
}

// OrderSchema is the schema response
type OrderSchema struct {
	ID        int64                `json:"id"`
	UUID      string               `json:"uuid"`
	Time      string               `json:"time"`
	InWord    int64                `json:"in_word"`
	Status    string               `json:"status"`
	UserID    int64                `json:"user_id"`
	Username  string               `json:"user_name"`
	UserEmail string               `json:"user_email"`
	Products  []ProductOrderSchema `json:"products"`
}

// CreateOrderSchema is the schema for creating an order request
type CreateOrderSchema struct {
	UserID   int64 `json:"user_id"`
	Products []struct {
		InventoryID int64 `json:"inventory_id"`
		Quantity    int64 `json:"quantity"`
	} `json:"product"`
}
