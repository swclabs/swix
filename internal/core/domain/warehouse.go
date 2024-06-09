package domain

import "github.com/shopspring/decimal"

type Warehouse struct {
	Id           string          `json:"id" gorm:"column:id"`
	ProductID    string          `json:"product_id" gorm:"column:product_id"`
	Model        string          `json:"model" gorm:"column:model"`
	Specs        string          `json:"specs" gorm:"column:specs"`
	Available    string          `json:"available" gorm:"column:available"`
	CurrencyCode string          `json:"currency_code" gorm:"column:currency_code"`
	Price        decimal.Decimal `json:"price" gorm:"column:price;type:decimal(19,4)"`
}

/*****************************************************************************/

type WarehouseSchemaReq struct {
	Id           string `json:"id" gorm:"column:id"`
	ProductID    string `json:"product_id" gorm:"column:product_id"`
	Price        string `json:"price" gorm:"column:price;type:decimal(19,4)"`
	Model        string `json:"model" gorm:"column:model"`
	Specs        string `json:"specs" gorm:"column:specs"`
	Available    string `json:"available" gorm:"column:available"`
	CurrencyCode string `json:"currency_code" gorm:"column:currency_code"`
}

type SpecsDetail struct {
	Color      string `json:"color"`
	Ram        string `json:"ram"`
	Ssd        string `json:"ssd"`
	ColorImage string `json:"color_image"`
	Image      string `json:"image"`
}

type WarehouseStruct struct {
	ProductID    string      `json:"product_id" validate:"required"`
	Price        string      `json:"price" validate:"number,required"`
	Model        string      `json:"model" validate:"required"`
	Available    string      `json:"available" validate:"number,required"`
	CurrencyCode string      `json:"currency_code" validate:"required"`
	Specs        SpecsDetail `json:"specs"`
}

type WarehouseSchema struct {
	Id string `json:"id"`
	WarehouseStruct
}
