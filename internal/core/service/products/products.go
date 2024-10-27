package products

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/core/domain/dtos"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/internal/core/domain/enum"
	"swclabs/swipex/internal/core/repos/categories"
	"swclabs/swipex/internal/core/repos/inventories"
	"swclabs/swipex/internal/core/repos/products"
	"swclabs/swipex/internal/core/repos/stars"
	"swclabs/swipex/pkg/infra/blob"
	"swclabs/swipex/pkg/infra/db"

	"github.com/jackc/pgx/v5"
	"github.com/shopspring/decimal"
)

var _ IProducts = (*Products)(nil)
var _ = app.Service(New)

// New creates a new ProductService object
func New(
	blob blob.IBlobStorage,
	products products.IProducts,
	inventory inventories.IInventories,
	category categories.ICategories,
	star stars.IStar,
) IProducts {
	return &Products{
		Blob:      blob,
		Products:  products,
		Inventory: inventory,
		Category:  category,
		Star:      star,
	}
}

// Products struct for product service
type Products struct {
	Blob      blob.IBlobStorage
	Products  products.IProducts
	Inventory inventories.IInventories
	Category  categories.ICategories
	Star      stars.IStar
}

// CreateProduct implements IProductService.
func (s *Products) CreateProduct(ctx context.Context, products dtos.Product) (int64, error) {
	_category, err := s.Category.GetByID(ctx, products.CategoryID)
	if err != nil {
		return -1, fmt.Errorf("category not found %v", err)
	}

	var types enum.Category
	if err := types.Load(_category.Name); err != nil {
		return -1, fmt.Errorf("category invalid %v", err)
	}
	var prd = entity.Product{
		Price:       products.Price,
		Description: products.Description,
		Name:        products.Name,
		SupplierID:  products.SupplierID,
		CategoryID:  products.CategoryID,
		Status:      products.Status,
		Specs:       "{}",
	}
	if products.Specs != nil {
		var specs, ok = products.Specs.(dtos.ProductSpecs)
		if !ok {
			return -1, fmt.Errorf("[code: %d] invalid specifications", http.StatusBadRequest)
		}
		specsByte, _ := json.Marshal(specs)
		prd.Specs = string(specsByte)
	}
	return s.Products.Insert(ctx, prd)
}

// InsertItem implements IProductService.
func (s *Products) InsertItem(ctx context.Context, product dtos.Inventory) error {
	var (
		price, _  = decimal.NewFromString(product.Price)
		avai, _   = strconv.Atoi(product.Available)
		inventory = entity.Inventory{
			Color:        product.Color,
			ColorImg:     product.ColorImg,
			ProductID:    product.ProductID,
			Price:        price,
			Available:    int64(avai),
			CurrencyCode: product.CurrencyCode,
			Status:       "active",
		}
	)
	items, err := s.Inventory.GetByColor(ctx, product.ProductID, product.Color)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}
	if len(items) > 0 {
		inventory.Image = items[0].Image
		inventory.ColorImg = items[0].ColorImg
	}
	tx, err := db.NewTx(ctx)
	if err != nil {
		return err
	}
	var invRepo = inventories.New(tx)
	specs, _ := json.Marshal(product.Specs)
	inventory.Specs = string(specs)
	_, err = invRepo.InsertProduct(ctx, inventory)
	if err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			return errTx
		}
		return err
	}
	return tx.Commit(ctx)
}
