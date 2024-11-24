package products

import "context"

// DeleteItem implements IProductService.
func (p *Products) DeleteItem(ctx context.Context, inventoryID int64) error {
	return p.Inventory.DeleteByID(ctx, inventoryID)
}

// DelProduct implements IProductService.
func (p *Products) DelProduct(ctx context.Context, productID int64) error {
	return p.Products.DeleteByID(ctx, productID)
}
