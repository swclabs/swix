package domain

// Categories Table
type Categories struct {
	Id          string `json:"id" gorm:"column:id"`
	Name        string `json:"name" gorm:"column:name" validate:"required"`
	Description string `json:"description" gorm:"column:description" validate:"required"`
}

/*****************************************************************************/

type CategorySlices struct {
	Data []Categories `json:"data"`
}

type CategoriesSwagger struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
