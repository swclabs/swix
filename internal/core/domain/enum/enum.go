// Package enum contains the enumerations used in the application.
package enum

import (
	"fmt"
)

// Category is an enumeration of the product categories.
type Category int

const (
	// Phone is a category of the product.
	Phone Category = 1 << iota

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

	// Watch is a category of the product
	Watch
)

// Storage is an enumeration of the electronic devices.
const (
	Storage   = Phone | Laptop | Tablet | Computer
	Accessory = Accessories | Earphone
)

// String returns the string representation of the Category.
func (c Category) String() string {
	switch {
	case c&Phone != 0:
		return "phone"
	case c&Tablet != 0:
		return "tablet"
	case c&Laptop != 0:
		return "laptop"
	case c&Earphone != 0:
		return "earphone"
	case c&Computer != 0:
		return "computer"
	case c&Accessories != 0:
		return "accessories"
	}
	return "unknown_category"
}

// Load loads the category type.
func (c *Category) Load(types string) error {
	switch types {
	case "phone":
		*c = Phone
	case "tablet":
		*c = Tablet
	case "laptop":
		*c = Laptop
	case "earphone":
		*c = Earphone
	case "computer":
		*c = Computer
	case "accessories":
		*c = Accessories
	default:
		return fmt.Errorf("invalid category type")
	}
	return nil
}
