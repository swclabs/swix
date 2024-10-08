package dtos

// DetailStorage is a type use to accept request and response
type DetailStorage struct {
	// Value 8GB
	RAM string `json:"RAM"`

	// Value 128GB
	SSD string `json:"SSD"`

	// Price 1.000.000 VND
	Price string `json:"price"`
}

// DetailWireless is a type use to accept request and response
type DetailWireless struct {
	Name  string `json:"name"`
	Desc  string `json:"desc"`
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

	Specs []SpecsItem `json:"specs"`
}

// ProductDetail is a type use to accept request and response
// T is a specifcation type
type ProductDetail struct {
	// Name of product
	Name string `json:"name"`

	// Screen 6.1 inch
	Screen string `json:"screen"`

	// Display Super AMOLED
	Display string `json:"display"`

	// Image of product
	Image []string `json:"image"`

	// Color of product
	Color []DetailColor `json:"color"`
}

// AccessoryDetail is a type use to accept request and response
type AccessoryDetail struct {
	Name     string   `json:"name"`
	Price    string   `json:"price"`
	Status   string   `json:"status"`
	Image    []string `json:"image"`
	Category string   `json:"category"`
}
