package model

// InventorySpecification is a type use to bind json in specs field of inventories table
type InventorySpecification struct {
	RAM string `json:"ram"`
	SSD string `json:"ssd"`
}
