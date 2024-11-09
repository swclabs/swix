// Package entity Addresses entities
package entity

// Address Table
type Address struct {
	ID       int64  `json:"id" db:"id"`
	UserID   int64  `json:"user_id" db:"user_id"`
	City     string `json:"city" db:"city"`
	Ward     string `json:"ward" db:"ward"`
	District string `json:"district" db:"district"`
	Street   string `json:"street" db:"street"`
}

type Province struct {
	PID  int64  `json:"pid" db:"pid"`
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type District struct {
	PID        int64  `json:"pid" db:"pid"`
	ID         string `json:"id" db:"id"`
	ProvinceID string `json:"province_id" db:"province_id"`
	Name       string `json:"name" db:"name"`
}

type Commune struct {
	PID        int64  `json:"pid" db:"pid"`
	ID         string `json:"id" db:"id"`
	DistrictID string `json:"district_id" db:"district_id"`
	Name       string `json:"name" db:"name"`
}
