// Package enum contains the enumerations used in the application.
package enum

import "fmt"

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

// ElectronicDevice is an enumeration of the electronic devices.
const ElectronicDevice = Phone | Laptop | Tablet | Computer

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
		"phone",
		"tablet",
		"laptop",
		"earphone",
		"computer",
		"accessories",
	}[c]
}

// Load loads the category type.
func (c Category) Load(types string) error {
	switch types {
	case "phone":
		c = Phone
	case "tablet":
		c = Tablet
	case "laptop":
		c = Laptop
	case "earphone":
		c = Earphone
	case "computer":
		c = Computer
	case "accessories":
		c = Accessories
	default:
		return fmt.Errorf("invalid category type")
	}
	return nil
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
