package domain

// Product : Table products
type Product struct {
	ID          int64  `json:"id" gorm:"column:id"`
	Image       string `json:"image" gorm:"column:image"`
	Price       string `json:"price" gorm:"column:price"`
	Description string `json:"description" gorm:"column:description"`
	Name        string `json:"name" gorm:"column:name"`
	SupplierID  int64  `json:"supplier_id" gorm:"column:supplier_id"`
	CategoryID  int64  `json:"category_id" gorm:"column:category_id"`
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

type IProductInCartRepository interface {
	GetByCartID(cartID int64) ([]ProductInCart, error)
	AddProduct(product *ProductInCart) error
	RemoveProduct(productID, cartID int64) error
	Save(product *ProductInCart) error
}
