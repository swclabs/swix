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
  "product_in_warehouse_id" bigint NOT NULL,
  "quantity" bigint NOT NULL,
  "total_amount" NUMERIC(19, 4) NOT NULL,
  "currency_code" varchar(3) NOT NULL
);
*/

// Orders table schema
type Orders struct {
	Id     int64  `json:"id" gorm:"column:id"`
	Uuid   string `json:"uuid" gorm:"column:uuid"`
	Time   string `json:"time" gorm:"column:time"`
	UserId int64  `json:"user_id" gorm:"column:user_id"`
	InWord int64  `json:"in_word" gorm:"column:in_word"`
	Status string `json:"status" gorm:"column:status"`
}

type ProductInOrder struct {
	Id                   int64           `json:"id" gorm:"column:id"`
	OrderId              int64           `json:"order_id" gorm:"column:order_id"`
	ProductInWarehouseId int64           `json:"product_in_warehouse_id" gorm:"column:product_in_warehouse_id"`
	Quantity             int64           `json:"quantity" gorm:"column:quantity"`
	CurrencyCode         string          `json:"currency_code" gorm:"column:currency_code"`
	TotalAmount          decimal.Decimal `json:"total_amount" gorm:"column:total_amount;type:decimal(19,4)"`
}

type ProductOrderSchema struct {
	Id                   int64  `json:"id"`
	OrderId              int64  `json:"order_id"`
	ProductInWarehouseId int64  `json:"product_in_warehouse_id" gorm:"column:product_in_warehouse_id"`
	Quantity             int64  `json:"quantity"`
	CurrencyCode         string `json:"currency_code"`
	TotalAmount          string `json:"total_amount"`
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
