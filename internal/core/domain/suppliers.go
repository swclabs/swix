package domain

const SuppliersTable = "suppliers"
const SuppliersAddressTable = "supplier_address"

// Suppliers table
type Suppliers struct {
	Id    string `json:"id" gorm:"column:id"`
	Name  string `json:"name" gorm:"column:name"`
	Email string `json:"email" gorm:"column:email"`
}

// SuppliersAddress suppliers address table
// CREATE TABLE "supplier_address" (
//
//		"supplier_id" bigint PRIMARY KEY,
//		"address_uuid" bigint NOT NULL
//	  );
type SuppliersAddress struct {
	SuppliersID string `json:"suppliers_id" gorm:"column:suppliers_id"`
	AddressUuiD string `json:"address_uuid" gorm:"column:address_uuid"`
}

type SuppliersListRes struct {
	Data []Suppliers `json:"data"`
}

type SuppliersReq struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"email,required"`
	City        string `json:"city" gorm:"column:city"`
	Ward        string `json:"ward" gorm:"column:ward"`
	District    string `json:"district" gorm:"column:district"`
	Street      string `json:"street" gorm:"column:street"`
}
