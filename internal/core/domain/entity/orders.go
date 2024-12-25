package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

// Order table schema
type Order struct {
	ID            int64           `json:"id" db:"id"`
	UUID          string          `json:"uuid" db:"uuid"`
	UserID        int64           `json:"user_id" db:"user_id"`
	DeliveryID    int64           `json:"delivery_id" db:"delivery_id"`
	Status        string          `json:"status" db:"status"`
	Time          time.Time       `json:"time" db:"time"`
	TotalAmount   decimal.Decimal `json:"total_amount" db:"total_amount"`
	PaymentMethod string          `json:"payment_method" db:"payment_method"`
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
