package dtos

// Supplier request, response
type Supplier struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"email,required"`
	City     string `json:"city" db:"city"`
	Ward     string `json:"ward" db:"ward"`
	District string `json:"district" db:"district"`
	Street   string `json:"street" db:"street"`
}
