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
	SupplierID  int64        `json:"supplier_id" validate:"number,required"`
	CategoryID  int64        `json:"category_id" validate:"number,required"`
	Status      string       `json:"status"`
}

// Product request, response
type Product struct {
	Specs       interface{} `json:"specs"`
	Price       string      `json:"price" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Name        string      `json:"name" validate:"required"`
	SupplierID  int64       `json:"supplier_id" validate:"number,required"`
	CategoryID  int64       `json:"category_id" validate:"number,required"`
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
	Category    string   `json:"category"`
}

// UpdateProductInfo request, response
type UpdateProductInfo struct {
	ID          int64        `json:"id" validate:"number,required"`
	Price       string       `json:"price"`
	Description string       `json:"description"`
	Name        string       `json:"name"`
	SupplierID  int64        `json:"supplier_id" validate:"number,omitempty"`
	CategoryID  int64        `json:"category_id" validate:"number,omitempty"`
	Status      string       `json:"status"`
	Specs       ProductSpecs `json:"specs"`
}

// CreateProduct response, request
type CreateProduct struct {
	Msg string `json:"msg"`
	ID  int64  `json:"id"`
}

// InvSpecification request, response
type InvSpecification struct {
	ID  int64  `json:"id"`
	RAM string `json:"ram"`
	SSD string `json:"ssd"`
}

// InvDetail request, response
type InvDetail struct {
	ProductID    string             `json:"product_id" validate:"number,required"`
	Price        string             `json:"price" validate:"number,required"`
	Available    string             `json:"available" validate:"number,required"`
	CurrencyCode string             `json:"currency_code" validate:"required"`
	ColorImg     string             `json:"color_img"`
	Color        string             `json:"color"`
	Status       string             `json:"status"`
	Image        []string           `json:"image"`
	Specs        []InvSpecification `json:"specs"`
}

// Inventory response, request
type Inventory struct {
	ID           int64              `json:"id"`
	ProductName  string             `json:"product_name"`
	ProductID    string             `json:"product_id" validate:"number,required"`
	Price        string             `json:"price" validate:"number,required"`
	Available    string             `json:"available" validate:"number,required"`
	CurrencyCode string             `json:"currency_code" validate:"required"`
	Category     string             `json:"category"`
	ColorImg     string             `json:"color_img"`
	Color        string             `json:"color"`
	Status       string             `json:"status"`
	Image        []string           `json:"image"`
	Specs        []InvSpecification `json:"specs"`
}

// StockHeader response, request
type StockHeader struct {
	All     int `json:"all"`
	Active  int `json:"active"`
	Draft   int `json:"draft"`
	Archive int `json:"archive"`
}

// InvStock response, request
type InvStock struct {
	Page   int         `json:"page"`
	Limit  int         `json:"limit"`
	Header StockHeader `json:"header"`
	Stock  []Inventory `json:"stock"`
}

// InvUpdate request, response
type InvUpdate struct {
	ID           string `json:"id" validate:"number,required"`
	ProductID    string `json:"product_id" validate:"omitempty,number"`
	Price        string `json:"price" validate:"omitempty,number"`
	Available    string `json:"available" validate:"omitempty,number"`
	CurrencyCode string `json:"currency_code"`
	Status       string `json:"status"`
}

// ProductView request, response
type ProductView struct {
	ID       int64       `json:"id"`
	Name     string      `json:"name"`
	Price    string      `json:"price"`
	Desc     string      `json:"desc"`
	Image    string      `json:"image"`
	Category string      `json:"category"`
	Specs    interface{} `json:"specs"`
}

// Specifications request, response
type Specifications struct {
	InventoryID int64  `json:"inventory_id" validate:"number,required"`
	RAM         string `json:"ram" validate:"number,required"`
	SSD         string `json:"ssd" validate:"number,required"`
}
