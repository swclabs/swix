package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

// Orders table schema
type Orders struct {
	ID          int64           `json:"id" db:"id"`
	UUID        string          `json:"uuid" db:"uuid"`
	UserID      int64           `json:"user_id" db:"user_id"`
	DeliveryID  int64           `json:"delivery_id" db:"delivery_id"`
	Status      string          `json:"status" db:"status"`
	Time        time.Time       `json:"time" db:"time"`
	TotalAmount decimal.Decimal `json:"total_amount" db:"total_amount"`
}

// ProductInOrder table schema
type ProductInOrder struct {
	ID           int64           `json:"id" db:"id"`
	OrderID      int64           `json:"order_id" db:"order_id"`
	InventoryID  int64           `json:"inventory_id" db:"inventory_id"`
	SpecsID      int64           `json:"specs_id" db:"specs_id"`
	Quantity     int64           `json:"quantity" db:"quantity"`
	CurrencyCode string          `json:"currency_code" db:"currency_code"`
	TotalAmount  decimal.Decimal `json:"total_amount" db:"total_amount"`
}
