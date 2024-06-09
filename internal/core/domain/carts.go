package domain

/**
CREATE TABLE "carts" (
  "id" bigserial PRIMARY KEY,
  "total_quantity" bigint NOT NULL,
  "user_id" bigint NOT NULL
);

CREATE TABLE "product_in_cart" (
  "id" bigserial PRIMARY KEY,
  "cart_id" bigint NOT NULL,
  "product_in_warehouse_id" bigint NOT NULL,
  "quantity" bigint NOT NULL
);
*/

// Carts table
type Carts struct {
	Id     int64 `json:"id" gorm:"column:id"`
	UserId int64 `json:"user_id" gorm:"column:user_id"`
}

// ProductInCart Table
type ProductInCart struct {
	Id                   int64 `json:"id" gorm:"column:id"`
	CartId               int64 `json:"cart_id" gorm:"column:cart_id"`
	ProductInWarehouseId int64 `json:"product_in_warehouse_id" gorm:"column:product_in_warehouse_id"`
	Quantity             int64 `json:"quantity" gorm:"column:quantity"`
}

type CartBodySchema struct {
	Img         string `json:"img"`
	ProductName string `json:"product_name"`
	Amount      string `json:"amount" gorm:"column:amount"`
	Quantity    int    `json:"quantity"`
	Category    string `json:"category" gorm:"column:category"`
}

// CartSchema schema
type CartSchema struct {
	Id       int64            `json:"id"`
	UserId   int64            `json:"user_id"`
	Products []CartBodySchema `json:"products"`
}
