package domain

// Suppliers table
type Suppliers struct {
	Id    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}

type SuppliersAddress struct {
	SuppliersID string `json:"suppliers_id" db:"suppliers_id"`
	AddressUuiD string `json:"address_uuid" db:"address_uuid"`
}

/*****************************************************************************/

type SupplierSchema struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"email,required"`
	City     string `json:"city" db:"city"`
	Ward     string `json:"ward" db:"ward"`
	District string `json:"district" db:"district"`
	Street   string `json:"street" db:"street"`
}
