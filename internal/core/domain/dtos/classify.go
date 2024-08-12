package dtos

// Supplier request, response
type Supplier struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"email,required"`
}
