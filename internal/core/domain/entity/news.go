package entity

import "time"

// CREATE TABLE "news" (
// 	"id" bigserial PRIMARY KEY,
// 	"created" timestamp default (timezone('utc', now())),
// 	"category" int NOT NULL,
// 	"header" varchar,
// 	"body" jsonb
//   )

type News struct {
	ID       int64     `json:"id" db:"id"`
	Category string    `json:"category" db:"category"`
	Header   string    `json:"header" db:"header"`
	Body     string    `json:"body" db:"body"`
	Created  time.Time `json:"created" db:"created"`
}
