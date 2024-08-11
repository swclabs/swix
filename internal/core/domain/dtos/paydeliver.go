package dtos

type DeliveryAddress struct {
	UserID   int64  `json:"user_id" validate:"required"`
	City     string `json:"city" validate:"required"`
	Ward     string `json:"ward" validate:"required"`
	District string `json:"district" validate:"required"`
	Street   string `json:"street" validate:"required"`
}

type Address struct {
	ID       int64  `json:"id" validate:"required"`
	City     string `json:"city" validate:"required"`
	Ward     string `json:"ward" validate:"required"`
	District string `json:"district" validate:"required"`
	Street   string `json:"street" validate:"required"`
}

type DeliveryBody struct {
	AddressID    int64  `json:"address_id" validate:"required"`
	UserID       int64  `json:"user_id" validate:"required"`
	Status       string `json:"status" validate:"required"`
	Method       string `json:"method" validate:"required"`
	Note         string `json:"note"`
	SentDate     string `json:"sent_date" validate:"date,omitempty"`
	ReceivedDate string `json:"received_date" validate:"date,omitempty"`
}

type Delivery struct {
	ID           int64  `json:"id" db:"id"`
	AddressID    int64  `json:"address_id" db:"address_id"`
	UserID       int64  `json:"user_id" db:"user_id"`
	Status       string `json:"status" db:"status"`
	Method       string `json:"method" db:"method"`
	Note         string `json:"note" db:"note"`
	SentDate     string `json:"sent_date" db:"sent_date"`
	ReceivedDate string `json:"received_date" db:"received_date"`
}
