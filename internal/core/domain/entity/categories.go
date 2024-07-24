package entity

// Categories Table
type Categories struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name" validate:"required"`
	Description string `json:"description" db:"description" validate:"required"`
}
