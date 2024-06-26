package domain

import "github.com/shopspring/decimal"

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
  "warehouse_id" bigint NOT NULL,
  "quantity" bigint NOT NULL,
  "total_amount" NUMERIC(19, 4) NOT NULL,
  "currency_code" varchar(3) NOT NULL
);
*/

// Orders table schema
type Orders struct {
	Id     int64  `json:"id" db:"id"`
	Uuid   string `json:"uuid" db:"uuid"`
	Time   string `json:"time" db:"time"`
	UserId int64  `json:"user_id" db:"user_id"`
	InWord int64  `json:"in_word" db:"in_word"`
	Status string `json:"status" db:"status"`
}

type ProductInOrder struct {
	Id           int64           `json:"id" db:"id"`
	OrderId      int64           `json:"order_id" db:"order_id"`
	WarehouseId  int64           `json:"product_in_warehouse_id" db:"product_in_warehouse_id"`
	Quantity     int64           `json:"quantity" db:"quantity"`
	CurrencyCode string          `json:"currency_code" db:"currency_code"`
	TotalAmount  decimal.Decimal `json:"total_amount" db:"total_amount"`
}

type ProductOrderSchema struct {
	Id           int64  `json:"id"`
	OrderId      int64  `json:"order_id"`
	WarehouseId  int64  `json:"product_in_warehouse_id" db:"product_in_warehouse_id"`
	Quantity     int64  `json:"quantity"`
	CurrencyCode string `json:"currency_code"`
	TotalAmount  string `json:"total_amount"`
}
type OrderSchema struct {
	Id        int64                `json:"id"`
	Uuid      string               `json:"uuid"`
	Time      string               `json:"time"`
	InWord    int64                `json:"in_word"`
	Status    string               `json:"status"`
	UserId    int64                `json:"user_id"`
	Username  string               `json:"user_name"`
	UserEmail string               `json:"user_email"`
	Products  []ProductOrderSchema `json:"products"`
}
