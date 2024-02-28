package domain

// Addresses Table
type Addresses struct {
	ID         int64  `json:"id" gorm:"column:id"`
	City       string `json:"city" gorm:"column:city"`
	Ward       string `json:"ward" gorm:"column:ward"`
	District   string `json:"district" gorm:"column:district"`
	Street     string `json:"street" gorm:"column:street"`
	UserId     string `json:"user_id" gorm:"column:user_id"`
	SupplierId string `json:"supplier_id" gorm:"column:supplier_id"`
}
