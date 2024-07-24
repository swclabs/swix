package entity

// Suppliers table
type Suppliers struct {
	ID    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}

// SuppliersAddress table
type SuppliersAddress struct {
	SuppliersID string `json:"suppliers_id" db:"suppliers_id"`
	AddressUuiD string `json:"address_uuid" db:"address_uuid"`
}
