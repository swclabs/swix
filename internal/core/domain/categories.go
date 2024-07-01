package domain

// Categories Table
type Categories struct {
	Id          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name" validate:"required"`
	Description string `json:"description" db:"description" validate:"required"`
}

/*****************************************************************************/

type CategoriesSwagger struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
