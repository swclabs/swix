package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	UserID         int64           `json:"user_id" db:"user_id"`
	UUID           string          `json:"uuid" db:"uuid"`
	Name           string          `json:"name" db:"name"`
	CategoryID     int64           `json:"category_id" db:"category_id"`
	Quantity       int             `json:"quantity" db:"quantity"`
	CurrencyCode   string          `json:"currency_code" db:"currency_code"`
	Color          string          `json:"color" db:"color"`
	InventoryImage string          `json:"inventory_image" db:"inventory_image"`
	Content        string          `json:"content" db:"content"`
	Time           time.Time       `json:"time" db:"time"`
	TotalAmount    decimal.Decimal `json:"total_amount" db:"total_amount"`
}
