package entity

import "time"

type Deliveries struct {
	ID           int64     `json:"id" db:"id"`
	AddressID    int64     `json:"address_id" db:"address_id"`
	UserID       int64     `json:"user_id" db:"user_id"`
	Status       string    `json:"status" db:"status"`
	Method       string    `json:"method" db:"method"`
	Note         string    `json:"note" db:"note"`
	SentDate     time.Time `json:"sent_date" db:"sent_date"`
	ReceivedDate time.Time `json:"received_date" db:"received_date"`
}
