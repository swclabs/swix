package domain

const WarehouseTable = "warehouse"

type Warehouse struct {
	Id        string `json:"id" gorm:"column:ram"`
	ProductID string `json:"product_id" gorm:"column:product_id"`
	Price     string `json:"price" gorm:"column:price"`
	Ram       string `json:"ram" gorm:"column:ram"`
	Ssd       string `json:"ssd" gorm:"column:ram"`
	Model     string `json:"model" gorm:"column:ram"`
	Available string `json:"available" gorm:"column:available"`
}
