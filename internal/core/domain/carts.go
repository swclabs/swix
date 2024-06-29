package domain

/**
CREATE TABLE "carts" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "inventory_id" bigint NOT NULL,
  "quantity" bigint NOT NULL
);
*/

// Carts table
type Carts struct {
	Id          int64 `json:"id" db:"id"`
	UserId      int64 `json:"user_id" db:"user_id"`
	InventoryId int64 `json:"inventory_id" db:"inventory_id"`
	Quantity    int64 `json:"quantity" db:"quantity"`
}
type CartBodySchema struct {
	Img         string `json:"img"`
	ProductName string `json:"product_name"`
	Amount      string `json:"amount" db:"amount"`
	Quantity    int64  `json:"quantity"`
	Category    string `json:"category" db:"category"`
}

// CartSchema schema
type CartSchema struct {
	UserId   int64            `json:"user_id"`
	Products []CartBodySchema `json:"products"`
}

type CartInsertReq struct {
	UserId      int64 `json:"user_id" validate:"required"`
	InventoryId int64 `json:"inventory_id" validate:"required"`
	Quantity    int64 `json:"quantity" validate:"required"`
}
