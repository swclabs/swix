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
	ID          int64 `json:"id" db:"id"`
	UserID      int64 `json:"user_id" db:"user_id"`
	InventoryID int64 `json:"inventory_id" db:"inventory_id"`
	Quantity    int64 `json:"quantity" db:"quantity"`
}

// CartSchema schema request, response
type CartSchema struct {
	Img         string `json:"img"`
	ProductName string `json:"product_name"`
	Amount      string `json:"amount" db:"amount"`
	Quantity    int64  `json:"quantity"`
	Category    string `json:"category" db:"category"`
}

// CartSlices schema
type CartSlices struct {
	UserID   int64        `json:"user_id"`
	Products []CartSchema `json:"products"`
}

// CartInsert request, response
type CartInsert struct {
	UserID      int64 `json:"user_id" validate:"required"`
	InventoryID int64 `json:"inventory_id" validate:"required"`
	Quantity    int64 `json:"quantity" validate:"required"`
}
