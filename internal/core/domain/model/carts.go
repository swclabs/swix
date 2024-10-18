package model

type Carts struct {
	Name           string `json:"name" db:"name"`
	CartID         int64  `json:"cart_id" db:"cart_id"`
	InventoryID    int64  `json:"inventory_id" db:"inventory_id"`
	ProductID      int64  `json:"product_id" db:"product_id"`
	Quantity       int64  `json:"quantity" db:"quantity"`
	Color          string `json:"color" db:"color"`
	InventoryPrice string `json:"inventory_price" db:"inventory_price"`
	CurrencyCode   string `json:"currency_code" db:"currency_code"`
	InventoryImage string `json:"inventory_image" db:"inventory_image"`
	InventorySpecs string `json:"inventory_specs" db:"inventory_specs"`
	CategoryName   string `json:"category_name" db:"category_name"`
}
