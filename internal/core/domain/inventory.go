package domain

import "github.com/shopspring/decimal"

type Inventories struct {
	Id           string          `json:"id" db:"id"`
	ProductID    int64           `json:"product_id" db:"product_id"`
	Model        string          `json:"model" db:"model"`
	Specs        string          `json:"specs" db:"specs"`
	Available    string          `json:"available" db:"available"`
	CurrencyCode string          `json:"currency_code" db:"currency_code"`
	Status       string          `json:"status" db:"status"`
	Price        decimal.Decimal `json:"price" db:"price"`
}

/*****************************************************************************/

type Inventory struct {
	Id           string `json:"id" db:"id"`
	ProductID    string `json:"product_id" db:"product_id"`
	Price        string `json:"price" db:"price"`
	Model        string `json:"model" db:"model"`
	Specs        string `json:"specs" db:"specs"`
	Available    string `json:"available" db:"available"`
	CurrencyCode string `json:"currency_code" db:"currency_code"`
}

type InventorySpecsDetail struct {
	Color      string `json:"color"`
	Ram        string `json:"ram"`
	Ssd        string `json:"ssd"`
	ColorImage string `json:"color_image"`
	Image      string `json:"image"`
}

type InventoryStruct struct {
	ProductID    string               `json:"product_id" validate:"required"`
	Price        string               `json:"price" validate:"number,required"`
	Model        string               `json:"model" validate:"required"`
	Available    string               `json:"available" validate:"number,required"`
	CurrencyCode string               `json:"currency_code" validate:"required"`
	Specs        InventorySpecsDetail `json:"specs"`
}

type InventorySchema struct {
	Id string `json:"id"`
	InventoryStruct
}

type InventoryStockSchema struct {
	All     int               `json:"all"`
	Active  int               `json:"active"`
	Draft   int               `json:"draft"`
	Archive int               `json:"archive"`
	Page    int               `json:"page"`
	Limit   int               `json:"limit"`
	Stock   []InventorySchema `json:"stock"`
}
