package model

// ProductXCategory is model of sql query join statement
// selectByCategory in products.sql.go
type ProductXCategory struct {
	ID           int64  `json:"id" db:"id"`
	Image        string `json:"image" db:"image"`
	Price        string `json:"price" db:"price"`
	Description  string `json:"description" db:"description"`
	Name         string `json:"name" db:"name"`
	Spec         string `json:"spec" db:"spec"`
	CategoryName string `json:"category_name" db:"category_name"`
}
