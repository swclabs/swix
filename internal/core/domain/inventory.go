package domain

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

/*****************************************************************************/

// Inventory table request, response
type Inventory struct {
	ID           string `json:"id" db:"id"`
	ProductID    string `json:"product_id" db:"product_id"`
	Price        string `json:"price" db:"price"`
	Specs        string `json:"specs" db:"specs"`
	Available    string `json:"available" db:"available"`
	CurrencyCode string `json:"currency_code" db:"currency_code"`
}

// InventorySpecsDetail request, response
type InventorySpecsDetail struct {
	Color      string   `json:"color"`
	RAM        string   `json:"ram"`
	Ssd        string   `json:"ssd"`
	ColorImage string   `json:"color_image"`
	Image      []string `json:"image"`
}

// InventoryStruct request, response
type InventoryStruct struct {
	ProductID    string               `json:"product_id" validate:"required"`
	Price        string               `json:"price" validate:"number,required"`
	Available    string               `json:"available" validate:"number,required"`
	CurrencyCode string               `json:"currency_code" validate:"required"`
	Specs        InventorySpecsDetail `json:"specs"`
}

// InventorySchema response, request
type InventorySchema struct {
	ID          string `json:"id"`
	ProductName string `json:"product_name"`
	Status      string `json:"status"`
	InventoryStruct
}

// InventoryStockHeader response, request
type InventoryStockHeader struct {
	All     int `json:"all"`
	Active  int `json:"active"`
	Draft   int `json:"draft"`
	Archive int `json:"archive"`
}

// InventoryStockSchema response, request
type InventoryStockSchema struct {
	Page   int                  `json:"page"`
	Limit  int                  `json:"limit"`
	Header InventoryStockHeader `json:"header"`
	Stock  []InventorySchema    `json:"stock"`
}

// InventoryDeviveSpecs request, response
type InventoryDeviveSpecs struct {
	ProductID string `json:"product_id"`
	RAM       string `json:"ram"`
	Ssd       string `json:"sd"`
	Color     string `json:"color"`
}
