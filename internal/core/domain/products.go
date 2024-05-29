package domain

const ProductsTable = "products"

// Products Table
type Products struct {
	ID          int64  `json:"id" gorm:"column:id"`
	Image       string `json:"image" gorm:"column:image"`
	Price       string `json:"price" gorm:"column:price"`
	Description string `json:"description" gorm:"column:description"`
	Name        string `json:"name" gorm:"column:name"`
	SupplierID  string `json:"supplier_id" gorm:"column:supplier_id"`
	CategoryID  string `json:"category_id" gorm:"column:category_id"`
	Spec        string `json:"spec" gorm:"column:spec"`
	Status      string `json:"status" gorm:"column:status"`
	Created     string `json:"created" gorm:"column:created"`
}

// ProductInCart Table
type ProductInCart struct {
	ID        int64 `json:"id" gorm:"column:id"`
	CartID    int64 `json:"cart_id" gorm:"column:cart_id"`
	ProductID int64 `json:"product_id" gorm:"column:product_id"`
	Amount    int64 `json:"amount" gorm:"column:amount"`
}

// ProductInOrder Table
type ProductInOrder struct {
	ID        int64 `json:"id" gorm:"column:id"`
	OrderID   int64 `json:"order_id" gorm:"column:order_id"`
	ProductID int64 `json:"product_id" gorm:"column:product_id"`
	Amount    int64 `json:"amount" gorm:"column:amount"`
}

// FavoriteProduct Table
type FavoriteProduct struct {
	ID        int64 `json:"id" gorm:"column:id"`
	UserID    int64 `json:"user_id" gorm:"column:user_id"`
	ProductID int64 `json:"product_id" gorm:"column:product_id"`
}

type Specs struct {
	Screen  string `json:"screen"`
	Display string `json:"display"`
	SSD     []int  `json:"SSD"`
	RAM     []int  `json:"RAM"`
}

type ProductReq struct {
	Specs
	Price       string `json:"price" validate:"required"`
	Description string `json:"description" validate:"required"`
	Name        string `json:"name" validate:"required"`
	SupplierID  string `json:"supplier_id" validate:"required"`
	CategoryID  string `json:"category_id" validate:"required"`
	Status      string `json:"status" validate:"required"`
}

type ProductRes struct {
	ID          int64    `json:"id"`
	Image       []string `json:"image"`
	Price       string   `json:"price"`
	Description string   `json:"description"`
	Name        string   `json:"name"`
	Status      string   `json:"status"`
	Created     string   `json:"created"`
	Spec        Specs    `json:"spec"`
}

type UploadProductRes struct {
	Msg string `json:"msg"`
	Id  int64  `json:"id"`
}

type ProductsRes struct {
	Data []ProductRes `json:"data" gorm:"column:data"`
}
