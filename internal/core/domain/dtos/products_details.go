package dtos

// DetailSSD is a type use to accept request and response
type DetailSSD struct {
	// Value 128GB
	Value string `json:"value"`

	// Price 1.000.000 VND
	Price string `json:"price"`
}

// DetailColor is a type use to accept request and response
type DetailColor struct {
	// Name Nature Titanium
	Name string `json:"name"`

	// Img of color Nature Titanium
	Img string `json:"img"`

	// Img of product Nature Titanium
	Product []string `json:"product"`
}

// ProductDetail is a type use to accept request and response
type ProductDetail struct {
	// Name of product
	Name string `json:"name"`

	// Screen 6.1 inch
	Screen string `json:"screen"`

	// Display Super AMOLED
	Display string `json:"display"`

	// Image of product
	Image []string `json:"image"`

	SSD   []DetailSSD   `json:"SSD"`
	Color []DetailColor `json:"color"`
}

// AccessoryDetail is a type use to accept request and response
type AccessoryDetail struct {
}
