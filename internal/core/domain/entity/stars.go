package entity

type Star struct {
	ID        int64 `json:"id" db:"id"`
	ProductID int64 `json:"product_id" db:"product_id"`
	UserID    int64 `json:"user_id" db:"user_id"`
}
