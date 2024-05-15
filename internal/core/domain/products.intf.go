package domain

import "context"

type IProductInCartRepository interface {
	GetByCartID(cartID int64) ([]ProductInCart, error)
	AddProduct(product *ProductInCart) error
	RemoveProduct(productID, cartID int64) error
	Save(product *ProductInCart) error
}

type IProductRepository interface {
	Insert(ctx context.Context, prd *Products) (int64, error)
	GetLimit(ctx context.Context, limit int) ([]ProductRes, error)
	UploadNewImage(ctx context.Context, urlImg string, id int) error
}
