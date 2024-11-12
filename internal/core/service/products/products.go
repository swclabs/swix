package products

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/core/domain/dtos"
	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/internal/core/domain/enum"
	"github.com/swclabs/swipex/internal/core/repos/categories"
	"github.com/swclabs/swipex/internal/core/repos/favorite"
	"github.com/swclabs/swipex/internal/core/repos/inventories"
	"github.com/swclabs/swipex/internal/core/repos/products"
	"github.com/swclabs/swipex/internal/core/repos/stars"
	"github.com/swclabs/swipex/pkg/infra/blob"
	"github.com/swclabs/swipex/pkg/infra/db"

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
	favorite favorite.IFavorite,
) IProducts {
	return &Products{
		Blob:      blob,
		Products:  products,
		Inventory: inventory,
		Category:  category,
		Star:      star,
		Favorite:  favorite,
	}
}

// Products struct for product service
type Products struct {
	Blob      blob.IBlobStorage
	Products  products.IProducts
	Inventory inventories.IInventories
	Category  categories.ICategories
	Star      stars.IStar
	Favorite  favorite.IFavorite
}

// AddBookmark implements IProducts.
func (s *Products) AddBookmark(ctx context.Context, userID int64, inventoryID int64) error {

	fav, err := s.Favorite.GetByInventoryID(ctx, inventoryID, userID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return s.Favorite.Save(ctx, entity.Favorite{UserID: userID, InventoryID: inventoryID})
	}
	return s.Favorite.Delete(ctx, entity.Favorite{UserID: fav.UserID, InventoryID: fav.InventoryID})
}

// GetBookmarks implements IProducts.
func (s *Products) GetBookmarks(ctx context.Context, userID int64) ([]dtos.Bookmark, error) {

	favories, err := s.Favorite.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var bookmarks = []dtos.Bookmark{}
	for _, fav := range favories {

		inv, err := s.Inventory.GetByID(ctx, fav.InventoryID)
		if err != nil {
			return nil, err
		}

		prod, err := s.Products.GetByID(ctx, inv.ProductID)
		if err != nil {
			return nil, err
		}

		category, err := s.Category.GetByID(ctx, prod.CategoryID)
		if err != nil {
			return nil, err
		}

		var (
			pSpecs dtos.ProductSpecs
			specs  dtos.Specs
		)

		if err := json.Unmarshal([]byte(prod.Specs), &pSpecs); err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(inv.Specs), &specs); err != nil {
			return nil, err
		}

		bookmark := dtos.Bookmark{
			ProductID: prod.ID,
			Category:  category.Name,
			Name:      prod.Name,
			Screen:    pSpecs.Screen,
			Display:   pSpecs.Display,
			Price:     prod.Price,
			Rating:    prod.Rating,
			Image:     strings.Split(prod.ShopImage, ","),
			Color: dtos.BookmarkItem{
				ColorName:  inv.Color,
				ColorImage: inv.ColorImg,
				Images:     strings.Split(inv.Image, ","),
				Specs: dtos.SpecsItem{
					InventoryID: inv.ID,
					SSD:         specs.SSD,
					RAM:         specs.RAM,
					Desc:        specs.Desc,
					Connection:  specs.Connection,
					Price:       inv.Price.String(),
				},
			},
		}

		if fav.InventoryID == inv.ID {
			bookmark.Color.Specs.Favorite = true
		}

		bookmarks = append(bookmarks, bookmark)
	}
	return bookmarks, nil
}

// RemoveBookmark implements IProducts.
func (s *Products) RemoveBookmark(ctx context.Context, userID int64, inventoryID int64) error {
	return s.Favorite.Delete(ctx, entity.Favorite{UserID: userID, InventoryID: inventoryID})
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
