package model

import (
	"github.com/shopspring/decimal"
)

type Order struct {
	CategoryID   int64           `json:"category_id" db:"category_id"`
	Quantity     int             `json:"quantity" db:"quantity"`
	CurrencyCode string          `json:"currency_code" db:"currency_code"`
	Color        string          `json:"color" db:"color"`
	Image        string          `json:"image" db:"image"`
	Name         string          `json:"name" db:"name"`
	TotalAmount  decimal.Decimal `json:"total_amount" db:"total_amount"`
	ItemSpecs    string          `json:"item_specs" db:"item_specs"`
}
