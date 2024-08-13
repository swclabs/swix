package dtos

// DetailSpecs is a type use to accept request and response
type DetailSpecs struct {
	// Value 8GB
	RAM string `json:"RAM"`

	// Value 128GB
	SSD string `json:"SSD"`

	// Price 1.000.000 VND
	Price string `json:"price"`
}

// DetailConnection is a type use to accept request and response
type DetailConnection struct {
}

// DetailColor is a type use to accept request and response
type DetailColor[T any] struct {
	// Name Nature Titanium
	Name string `json:"name"`

	// Img of color Nature Titanium
	Img string `json:"img"`

	// Img of product Nature Titanium
	Product []string `json:"product"`

	Specs []T `json:"specs"`
}

// ProductDetail is a type use to accept request and response
// T is a specifcation type
type ProductDetail[T any] struct {
	// Name of product
	Name string `json:"name"`

	// Screen 6.1 inch
	Screen string `json:"screen"`

	// Display Super AMOLED
	Display string `json:"display"`

	// Image of product
	Image []string `json:"image"`

	// Color of product
	Color []DetailColor[T] `json:"color"`
}

// AccessoryDetail is a type use to accept request and response
type AccessoryDetail struct {
}
