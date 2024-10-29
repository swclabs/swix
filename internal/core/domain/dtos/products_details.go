package dtos

// Color is a type use to accept request and response
type Color struct {
	// Name Nature Titanium
	Name string `json:"name"`

	// ImageColor of color Nature Titanium
	ImageColor string `json:"img"`

	// Img of product Nature Titanium
	Product []string `json:"product"`

	Specs []SpecsItem `json:"specs"`
}

// ProductDetail is a type use to accept request and response
type ProductDetail struct {
	// Name of product
	Name string `json:"name"`

	// Screen 6.1 inch
	Screen string `json:"screen"`

	// Display Super AMOLED
	Display string `json:"display"`

	Price string `json:"price"`

	Rating float64 `json:"rating"`

	// Image of product
	Image []string `json:"image"`

	// Color of product
	Color []Color `json:"color"`
}
