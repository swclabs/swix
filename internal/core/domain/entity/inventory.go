package entity

import "github.com/shopspring/decimal"

// Inventory table
type Inventory struct {
	ID           int64           `json:"id" db:"id"`
	ProductID    int64           `json:"product_id" db:"product_id"`
	Available    int64           `json:"available" db:"available"`
	CurrencyCode string          `json:"currency_code" db:"currency_code"`
	Status       string          `json:"status" db:"status"`
	Color        string          `json:"color" db:"color"`
	ColorImg     string          `json:"color_img" db:"color_img"`
	Image        string          `json:"image" db:"image"`
	Specs        string          `json:"specs" db:"specs"`
	Price        decimal.Decimal `json:"price" db:"price"`
}
