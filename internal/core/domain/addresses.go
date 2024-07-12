package domain

// Addresses Table
type Addresses struct {
	ID       int64  `json:"id" db:"id"`
	UUID     string `json:"uuid" db:"uuid"`
	City     string `json:"city" db:"city"`
	Ward     string `json:"ward" db:"ward"`
	District string `json:"district" db:"district"`
	Street   string `json:"street" db:"street"`
}
