package domain

// Categories Table
type Categories struct {
	Id          string `json:"id" gorm:"column:id"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}

type CategoriesList struct {
	Data []Categories `json:"data"`
}

type CategoriesReq struct {
	Name        string `json:"name" validate:"required" gorm:"column:name"`
	Description string `json:"description" validate:"required" gorm:"column:description"`
}
