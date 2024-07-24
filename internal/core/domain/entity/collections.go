package entity

import "time"

// Collection use to query data from database, 
// define in database, used to save a CollectionBody
type Collection struct {
	ID       int64     `json:"id" db:"id"`
	Position string    `json:"position" db:"position"` // Example: mac#1, mac#2
	Headline string    `json:"headline" db:"headline"` // Ex: Get to know Mac
	Body     string    `json:"body" db:"body"`
	Created  time.Time `json:"created" db:"created"`
}