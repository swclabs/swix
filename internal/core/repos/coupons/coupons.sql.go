package coupons

// type Coupons struct {
// 	ID          int64     `json:"id" db:"id"`
// 	Code        string    `json:"code" db:"code"`
// 	Used        int       `json:"used" db:"used"`
// 	Status      string    `json:"status" db:"status"`
// 	MaxUse      int       `json:"max_use" db:"max_use"`
// 	Discount    int       `json:"discount" db:"discount"`
// 	ExpiredAt   time.Time `json:"expired_at" db:"expired_at"`
// 	Description string    `json:"description" db:"description"`
// }

// type CouponsUsed struct {
// 	ID       int64     `json:"id" db:"id"`
// 	UserID   int64     `json:"user_id" db:"user_id"`
// 	CouponCode int64     `json:"coupon_code" db:"coupon_code"`
// 	OrderID  int64     `json:"order_id" db:"order_id"`
// 	UsedAt   time.Time `json:"used_at" db:"used_at"`
// }

const (
	insert = `
		INSERT INTO coupons (code, discount, status, used, max_use, description, expired_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
	`

	useCoupons = `
		INSERT INTO coupons_used (user_id, coupon_code, order_id, used_at)
		VALUES ($1, $2, $3, $4);
	`

	getAll = `
		SELECT * FROM coupons;
	`

	getByUser = `
		SELECT * FROM coupons_used WHERE user_id = $1;
	`

	getByCode = `
		SELECT * FROM coupons WHERE code = $1;
	`

	delete = `
		DELETE FROM coupons WHERE code = $1;
	`
)
