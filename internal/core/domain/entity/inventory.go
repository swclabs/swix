package entity

import "github.com/shopspring/decimal"

// Inventories table
type Inventories struct {
	ID           string          `json:"id" db:"id"`
	ProductID    int64           `json:"product_id" db:"product_id"`
	Specs        string          `json:"specs" db:"specs"`
	Available    string          `json:"available" db:"available"`
	CurrencyCode string          `json:"currency_code" db:"currency_code"`
	Status       string          `json:"status" db:"status"`
	Price        decimal.Decimal `json:"price" db:"price"`
}
