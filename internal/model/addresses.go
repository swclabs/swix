package model

// Address :Table addresses
type Address struct {
	ID       int64  `json:"id" gorm:"column:id"`
	City     string `json:"city" gorm:"column:city"`
	Ward     string `json:"ward" gorm:"column:ward"`
	District string `json:"district" gorm:"column:district"`
	Street   string `json:"street" gorm:"column:street"`
}
