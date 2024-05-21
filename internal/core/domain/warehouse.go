package domain

const WarehouseTable = "warehouse"

type Warehouse struct {
	Id        string `json:"id" gorm:"column:id"`
	ProductID string `json:"product_id" gorm:"column:product_id"`
	Price     string `json:"price" gorm:"column:price"`
	Model     string `json:"model" gorm:"column:model"`
	Specs     string `json:"specs" gorm:"column:specs"`
	Available string `json:"available" gorm:"column:available"`
}

type SpecsDetail struct {
	Color      string `json:"color"`
	Ram        string `json:"ram"`
	Ssd        string `json:"ssd"`
	ColorImage string `json:"color_image"`
	Image      string `json:"image"`
}

type WarehouseReq struct {
	ProductID string      `json:"product_id" validate:"required"`
	Price     string      `json:"price" validate:"required"`
	Model     string      `json:"model" validate:"required"`
	Specs     SpecsDetail `json:"specs"`
	Available string      `json:"available" validate:"required"`
}

type WarehouseRes struct {
	Id string `json:"id"`
	WarehouseReq
}
