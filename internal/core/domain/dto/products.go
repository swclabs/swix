package dto

// Specs request, response
type Specs struct {
	Screen  string `json:"screen"`
	Display string `json:"display"`
	SSD     []int  `json:"SSD"`
	RAM     []int  `json:"RAM"`
}

// Product request, response
type Product struct {
	Specs
	Price       string `json:"price" validate:"required"`
	Description string `json:"description" validate:"required"`
	Name        string `json:"name" validate:"required"`
	SupplierID  string `json:"supplier_id" validate:"required"`
	CategoryID  string `json:"category_id" validate:"required"`
	Status      string `json:"status" validate:"required"`
}

// ProductSchema request, response
type ProductSchema struct {
	ID          int64    `json:"id"`
	Image       []string `json:"image"`
	Price       string   `json:"price"`
	Description string   `json:"description"`
	Name        string   `json:"name"`
	Status      string   `json:"status"`
	Created     string   `json:"created"`
	Spec        Specs    `json:"spec"`
}

// UpdateProductInfo request, response
type UpdateProductInfo struct {
	Product
	ID int64 `json:"id" validate:"required"`
}

// CreateProduct response, request
type CreateProduct struct {
	Msg string `json:"msg"`
	ID  int64  `json:"id"`
}

// InventorySpecsDetail request, response
type InventorySpecsDetail struct {
	Color      string   `json:"color"`
	RAM        string   `json:"ram"`
	Ssd        string   `json:"ssd"`
	ColorImage string   `json:"color_image"`
	Image      []string `json:"image"`
}

// InventoryDetail request, response
type InventoryDetail struct {
	ProductID    string               `json:"product_id" validate:"number,required"`
	Price        string               `json:"price" validate:"number,required"`
	Available    string               `json:"available" validate:"number,required"`
	CurrencyCode string               `json:"currency_code" validate:"required"`
	Specs        InventorySpecsDetail `json:"specs"`
}

// Inventory response, request
type Inventory struct {
	ID          string `json:"id"`
	ProductName string `json:"product_name"`
	Status      string `json:"status"`
	InventoryDetail
}

// StockHeader response, request
type StockHeader struct {
	All     int `json:"all"`
	Active  int `json:"active"`
	Draft   int `json:"draft"`
	Archive int `json:"archive"`
}

// StockInInventory response, request
type StockInInventory struct {
	Page   int         `json:"page"`
	Limit  int         `json:"limit"`
	Header StockHeader `json:"header"`
	Stock  []Inventory `json:"stock"`
}

// InventoryDeviceSpecs request, response
type InventoryDeviceSpecs struct {
	ProductID string `json:"product_id"`
	RAM       string `json:"ram"`
	Ssd       string `json:"ssd"`
	Color     string `json:"color"`
}

// Supplier request, response
type Supplier struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"email,required"`
	City     string `json:"city" db:"city"`
	Ward     string `json:"ward" db:"ward"`
	District string `json:"district" db:"district"`
	Street   string `json:"street" db:"street"`
}

// UpdateInventory request, response
type UpdateInventory struct {
	ID           string               `json:"id" validate:"number"`
	ProductID    string               `json:"product_id" validate:"number"`
	Price        string               `json:"price" validate:"number"`
	Available    string               `json:"available" validate:"number"`
	CurrencyCode string               `json:"currency_code"`
	Status       string               `json:"status"`
	Specs        InventorySpecsDetail `json:"specs"`
}
