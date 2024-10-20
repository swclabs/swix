package dtos

// DeliveryAddress request, response
type DeliveryAddress struct {
	UserID   int64  `json:"user_id" validate:"required"`
	City     string `json:"city" validate:"required"`
	Ward     string `json:"ward" validate:"required"`
	District string `json:"district" validate:"required"`
	Street   string `json:"street" validate:"required"`
}

// Address request, response
type Address struct {
	ID       int64  `json:"id" validate:"required"`
	City     string `json:"city" validate:"required"`
	Ward     string `json:"ward" validate:"required"`
	District string `json:"district" validate:"required"`
	Street   string `json:"street" validate:"required"`
}

// DeliveryBody request, response
type DeliveryBody struct {
	AddressID    int64  `json:"address_id" validate:"required"`
	UserID       int64  `json:"user_id" validate:"required"`
	Status       string `json:"status" validate:"required"`
	Method       string `json:"method" validate:"required"`
	Note         string `json:"note"`
	SentDate     string `json:"sent_date" validate:"date,omitempty"`
	ReceivedDate string `json:"received_date" validate:"date,omitempty"`
}

// Delivery request, response
type Delivery struct {
	ID           int64   `json:"id" db:"id"`
	Address      Address `json:"address" db:"address"`
	UserID       int64   `json:"user_id" db:"user_id"`
	Status       string  `json:"status" db:"status"`
	Method       string  `json:"method" db:"method"`
	Note         string  `json:"note" db:"note"`
	SentDate     string  `json:"sent_date" db:"sent_date"`
	ReceivedDate string  `json:"received_date" db:"received_date"`
}
