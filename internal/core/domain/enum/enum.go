// Package enum contains the enumerations used in the application.
package enum

// Category is an enumeration of the product categories.
type Category int

// Accessory is an enumeration of the product accessories.
type Accessory int

const (
	// Phone is a category of the product.
	Phone Category = iota

	// Tablet is a category of the product.
	Tablet

	// Laptop is a category of the product.
	Laptop

	// Earphone is a category of the product.
	Earphone

	// Computer is a category of the product.
	Computer

	// Accessories is a category of the product.
	Accessories
)

const (
	// PhoneAccessory is an accessory of the product.
	PhoneAccessory Accessory = iota

	// TabletAccessory is an accessory of the product.
	TabletAccessory

	// LaptopAccessory is an accessory of the product.
	LaptopAccessory

	// EarphoneAccessory is an accessory of the product.
	EarphoneAccessory

	// ComputerAccessory is an accessory of the product.
	ComputerAccessory
)

// String returns the string representation of the Category.
func (c Category) String() string {
	return [...]string{
		"iPhone",
		"iPad",
		"Macbook",
		"AirPod",
		"iMac",
		"Accessories",
	}[c]
}

// String returns the string representation of the Accessory.
func (a Accessory) String() string {
	return [...]string{
		"iPhone",
		"iPad",
		"Macbook",
		"AirPod",
		"iMac",
	}[a]
}
