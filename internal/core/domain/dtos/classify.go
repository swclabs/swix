package dtos

// Supplier request, response
type Supplier struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"email,required"`
}

type UpdateCategories struct {
	ID          int64  `json:"id" db:"id"`
	Name        string `json:"name" db:"name" validate:"required"`
	Description string `json:"description" db:"description" validate:"required"`
}
