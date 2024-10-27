package products

import "context"

// DeleteItem implements IProductService.
func (s *Products) DeleteItem(ctx context.Context, inventoryID int64) error {
	return s.Inventory.DeleteByID(ctx, inventoryID)
}

// DelProduct implements IProductService.
func (s *Products) DelProduct(ctx context.Context, productID int64) error {
	return s.Products.DeleteByID(ctx, productID)
}
