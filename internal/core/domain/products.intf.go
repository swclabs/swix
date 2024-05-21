package domain

import "context"

// IProductInCartRepository defines methods to interact with products in a shopping cart (ProductInCart) data.
type IProductInCartRepository interface {
	// GetByCartID retrieves a list of products in the cart specified by the cartID.
	// cartID is the ID of the shopping cart.
	// Returns a slice of ProductInCart objects and an error if any issues occur during the retrieval process.
	GetByCartID(cartID int64) ([]ProductInCart, error)

	// AddProduct adds a new product to the shopping cart.
	// product is a pointer to the ProductInCart object to be added.
	// Returns an error if any issues occur during the addition process.
	AddProduct(product *ProductInCart) error

	// RemoveProduct removes a product from the shopping cart.
	// productID is the ID of the product to be removed.
	// cartID is the ID of the shopping cart from which the product will be removed.
	// Returns an error if any issues occur during the removal process.
	RemoveProduct(productID, cartID int64) error

	// Save updates the information of an existing product in the shopping cart.
	// product is a pointer to the ProductInCart object to be saved.
	// Returns an error if any issues occur during the saving process.
	Save(product *ProductInCart) error
}

// IProductRepository defines methods to interact with product (Products) data.
type IProductRepository interface {
	// Insert adds a new product to the database.
	// ctx is the context to manage the request's lifecycle.
	// prd is a pointer to the Products object to be added.
	// Returns the ID of the newly inserted product and an error if any issues occur during the insertion process.
	Insert(ctx context.Context, prd *Products) (int64, error)

	// GetLimit retrieves a list of products with a specified limit.
	// ctx is the context to manage the request's lifecycle.
	// limit is the maximum number of products to retrieve.
	// Returns a slice of ProductRes objects and an error if any issues occur during the retrieval process.
	GetLimit(ctx context.Context, limit int) ([]ProductRes, error)

	// UploadNewImage updates the image URL of a specified product.
	// ctx is the context to manage the request's lifecycle.
	// urlImg is the new image URL to be uploaded.
	// id is the ID of the product to be updated.
	// Returns an error if any issues occur during the update process.
	UploadNewImage(ctx context.Context, urlImg string, id int) error
}
