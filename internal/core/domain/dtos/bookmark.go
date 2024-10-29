package dtos

type Bookmark struct {
	ProductID int64 `json:"product_id"`

	Category string `json:"category"`

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
	Color BookmarkItem `json:"color"`
}

type BookmarkItem struct {
	// ColorName Nature Titanium
	ColorName string `json:"color_name"`

	// ColorImage of color Nature Titanium
	ColorImage string `json:"color_img"`

	// Img of product Nature Titanium
	Images []string `json:"images"`

	Specs SpecsItem `json:"specs"`
}
