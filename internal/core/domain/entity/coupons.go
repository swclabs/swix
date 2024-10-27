package entity

import "time"

// CREATE TABLE "coupons" (
// 	"id" bigserial PRIMARY KEY,
// 	"code" varchar UNIQUE NOT NULL,
// 	"discount" NUMERIC(19, 4) NOT NULL,
// 	"status" varchar NOT NULL,
// 	"used" int NOT NULL,
// 	"max_use" int NOT NULL,
// 	"description" varchar NOT NULL,
// 	"expired_at" timestamptz NOT NULL
//   );

type Coupons struct {
	ID          int64     `json:"id" db:"id"`
	Code        string    `json:"code" db:"code"`
	Used        int       `json:"used" db:"used"`
	Status      string    `json:"status" db:"status"`
	MaxUse      int       `json:"max_use" db:"max_use"`
	Discount    int       `json:"discount" db:"discount"`
	ExpiredAt   time.Time `json:"expired_at" db:"expired_at"`
	Description string    `json:"description" db:"description"`
}

type CouponsUsed struct {
	ID         int64     `json:"id" db:"id"`
	UserID     int64     `json:"user_id" db:"user_id"`
	CouponCode string    `json:"coupon_code" db:"coupon_code"`
	OrderID    int64     `json:"order_id" db:"order_id"`
	UsedAt     time.Time `json:"used_at" db:"used_at"`
}
