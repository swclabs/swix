package dtos

// ProductSpecs request, response
type ProductSpecs struct {
	Screen  string `json:"screen"`
	Display string `json:"display"`
}

// ProductRequest request, response
type ProductRequest struct {
	Specs       ProductSpecs `json:"specs"`
	Price       string       `json:"price" validate:"required"`
	Description string       `json:"description" validate:"required"`
	Name        string       `json:"name" validate:"required"`
	SupplierID  string       `json:"supplier_id" validate:"number,required"`
	CategoryID  string       `json:"category_id" validate:"number,required"`
	Status      string       `json:"status"`
}

// Product request, response
type Product struct {
	Specs       interface{} `json:"specs"`
	Price       string      `json:"price" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Name        string      `json:"name" validate:"required"`
	SupplierID  string      `json:"supplier_id" validate:"number,required"`
	CategoryID  string      `json:"category_id" validate:"number,required"`
	Status      string      `json:"status"`
}

// ProductResponse request, response
type ProductResponse struct {
	ID          int64    `json:"id"`
	Image       []string `json:"image"`
	Price       string   `json:"price"`
	Description string   `json:"description"`
	Name        string   `json:"name"`
	Status      string   `json:"status"`
	Created     string   `json:"created"`
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

// InventorySpecification request, response
type InventorySpecification struct {
	RAM string `json:"ram"`
	SSD string `json:"ssd"`
}

// InventoryDetail request, response
type InventoryDetail struct {
	ProductID    string                   `json:"product_id" validate:"number,required"`
	Price        string                   `json:"price" validate:"number,required"`
	Available    string                   `json:"available" validate:"number,required"`
	CurrencyCode string                   `json:"currency_code" validate:"required"`
	ColorImg     string                   `json:"color_img"`
	Color        string                   `json:"color"`
	Status       string                   `json:"status"`
	Image        []string                 `json:"image"`
	Specs        []InventorySpecification `json:"specs"`
}

// Inventory response, request
type Inventory struct {
	ID           string                   `json:"id"`
	ProductName  string                   `json:"product_name"`
	ProductID    string                   `json:"product_id" validate:"number,required"`
	Price        string                   `json:"price" validate:"number,required"`
	Available    string                   `json:"available" validate:"number,required"`
	CurrencyCode string                   `json:"currency_code" validate:"required"`
	ColorImg     string                   `json:"color_img"`
	Color        string                   `json:"color"`
	Status       string                   `json:"status"`
	Image        []string                 `json:"image"`
	Specs        []InventorySpecification `json:"specs"`
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

// UpdateInventory request, response
type UpdateInventory struct {
	ID           string `json:"id" validate:"number,require"`
	ProductID    string `json:"product_id" validate:"number,require"`
	Price        string `json:"price" validate:"number,require"`
	Available    string `json:"available" validate:"number,require"`
	CurrencyCode string `json:"currency_code"`
	Status       string `json:"status"`
}

// ProductView request, response
type ProductView struct {
	ID    int64       `json:"id"`
	Name  string      `json:"name"`
	Price string      `json:"price"`
	Desc  string      `json:"desc"`
	Image string      `json:"image"`
	Specs interface{} `json:"specs"`
}
