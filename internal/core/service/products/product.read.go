package products

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/swclabs/swipex/internal/core/domain/dtos"
	"github.com/swclabs/swipex/internal/core/domain/entity"
	"github.com/swclabs/swipex/internal/core/domain/enum"
	swcerr "github.com/swclabs/swipex/pkg/lib/errors"
	"github.com/swclabs/swipex/pkg/utils"

	"github.com/jackc/pgx/v5"
)

// GetProductInfo implements IProducts.
func (p *Products) GetProductInfo(ctx context.Context, productID int64) (*dtos.ProductResponse, error) {
	product, err := p.Products.GetByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	category, err := p.Category.GetByID(ctx, product.CategoryID)
	if err != nil {
		return nil, err
	}

	resp := dtos.ProductResponse{
		ID:          product.ID,
		Price:       product.Price,
		Description: product.Description,
		Name:        product.Name,
		Status:      product.Status,
		Image:       "",
		Created:     utils.HanoiTimezone(product.Created),
		Category:    category.Name,
	}


	if len(strings.Split(product.Image, ",")) > 0 {
		resp.Image = strings.Split(product.Image, ",")[0]
	}

	return &resp, nil
}

// SearchDetails implements IProducts.
func (p *Products) SearchDetails(ctx context.Context, userID int64, keyword string) ([]dtos.ProductDetail, error) {
	products, err := p.Products.Search(ctx, keyword)
	if err != nil {
		return nil, swcerr.Service("keyword error", err)
	}

	var details []dtos.ProductDetail
	for _, product := range products {

		detail, err := p.Detail(ctx, userID, product.ID)
		if err != nil {
			return nil, err
		}

		details = append(details, *detail)
	}

	return details, nil
}

// ProductType implements IProductService.
func (p *Products) ProductType(ctx context.Context, types enum.Category, offset int) ([]dtos.ProductDTO, error) {
	products, err := p.Products.GetByCategory(ctx, types, offset)
	if err != nil {
		return nil, err
	}

	var productView []dtos.ProductDTO
	for _, p := range products {

		_view := dtos.ProductDTO{
			ID:       p.ID,
			Price:    p.Price,
			Desc:     p.Description,
			Name:     p.Name,
			Image:    p.Image,
			Rating:   p.Rating,
			Category: p.CategoryName,
		}

		var specs dtos.ProductSpecs
		if err := json.Unmarshal([]byte(p.Specs), &specs); err != nil {
			return nil, fmt.Errorf("[code %d] %v", http.StatusBadRequest, err)
		}

		_view.Specs = specs
		productView = append(productView, _view)
	}

	return productView, nil
}

// GetItem implements IProductService.
func (p *Products) GetItem(ctx context.Context, inventoryID int64) (*dtos.Inventory, error) {
	item, err := p.Inventory.GetByID(ctx, inventoryID)
	if err != nil {
		return nil, err
	}

	product, err := p.Products.GetByID(ctx, item.ProductID)
	if err != nil {
		return nil, err
	}

	category, _ := p.Category.GetByID(ctx, product.CategoryID)
	var specs dtos.Specs
	if err := json.Unmarshal([]byte(item.Specs), &specs); err != nil {
		return nil, err
	}

	var (
		result = dtos.Inventory{
			ID:           item.ID,
			ProductName:  product.Name,
			ProductID:    item.ProductID,
			Price:        item.Price.String(),
			Available:    strconv.Itoa(int(item.Available)),
			CurrencyCode: item.CurrencyCode,
			Status:       item.Status,
			Color:        item.Color,
			ColorImg:     item.ColorImg,
			Image:        strings.Split(item.Image, ","),
			Category:     category.Name,
			Specs:        specs,
			ItemCode:     category.Name + "#" + strconv.Itoa(int(item.ID)),
		}
	)

	return &result, nil

}

// Detail implements IProductService.
func (p *Products) Detail(ctx context.Context, userID int64, productID int64) (*dtos.ProductDetail, error) {
	var (
		productSpecs dtos.ProductSpecs
		details      dtos.ProductDetail
	)

	colors, err := p.Inventory.GetColor(ctx, productID)
	if err != nil {
		return nil, err
	}

	product, err := p.Products.GetByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(product.Specs), &productSpecs); err != nil {
		return nil, err
	}

	details.Name = product.Name
	details.Screen = productSpecs.Screen
	details.Display = productSpecs.Display
	details.Price = product.Price
	details.Image = strings.Split(product.ShopImage, ",")
	details.Rating = product.Rating
	details.Color = []dtos.Color{}

	for _, color := range colors {
		items, err := p.Inventory.GetByColor(ctx, productID, color.Color)
		if err != nil {
			return nil, err
		}

		if len(items) == 0 {
			continue
		}

		detailsColor := dtos.Color{
			Name:       color.Color,
			ImageColor: items[0].ColorImg,
			Product:    strings.Split(items[0].Image, ","),
		}

		for _, item := range items {
			var spec dtos.SpecsItem
			if err := json.Unmarshal([]byte(item.Specs), &spec); err != nil {
				return nil, err
			}

			if userID != -1 {
				favorite, err := p.Favorite.GetByInventoryID(ctx, item.ID, userID)
				if err != nil && !errors.Is(err, pgx.ErrNoRows) {
					return nil, err
				}

				if favorite != nil && favorite.InventoryID == item.ID {
					spec.Favorite = true
				}
			}

			spec.Price = item.Price.String()
			spec.InventoryID = item.ID
			detailsColor.Specs = append(detailsColor.Specs, spec)
		}

		details.Color = append(details.Color, detailsColor)
	}
	return &details, nil
}

// GetInvItems implements IProductService.
func (p *Products) GetInvItems(ctx context.Context, page int, limit int) (*dtos.InventoryItems, error) {
	var invItems dtos.InventoryItems

	items, err := p.Inventory.GetLimit(ctx, limit, page)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		product, err := p.Products.GetByID(ctx, item.ProductID)
		if err != nil {
			return nil, err
		}

		category, err := p.Category.GetByID(ctx, product.CategoryID)
		if err != nil {
			return nil, err
		}

		switch item.Status {
		case "active":
			invItems.Header.Active++
		case "draft":
			invItems.Header.Draft++
		case "archived":
			invItems.Header.Archive++
		}

		invItems.Header.All++
		_item := dtos.Inventory{
			ID:           item.ID,
			ProductImg:   product.Image,
			ProductName:  product.Name,
			ProductID:    item.ProductID,
			Price:        item.Price.String(),
			Available:    strconv.Itoa(int(item.Available)),
			CurrencyCode: item.CurrencyCode,
			Status:       item.Status,
			Color:        item.Color,
			ColorImg:     item.ColorImg,
			Image:        strings.Split(item.Image, ","),
			Category:     category.Name,
			Specs:        dtos.Specs{},
			ItemCode:     strings.ToUpper(category.Name) + "#" + strconv.Itoa(int(item.ID)),
		}

		var specs dtos.Specs
		if err := json.Unmarshal([]byte(item.Specs), &specs); err != nil {
			return nil, err
		}

		_item.Specs = specs
		invItems.Stock = append(invItems.Stock, _item)
	}

	invItems.Page = page
	invItems.Limit = limit

	return &invItems, nil
}

// GetItems implements IProductService.
func (p *Products) GetItems(ctx context.Context, productID int64) ([]entity.Inventory, error) {
	return p.Inventory.GetByProductID(ctx, productID)
}

// Search implements IProductService.
func (p *Products) Search(ctx context.Context, keyword string) ([]dtos.ProductResponse, error) {
	_products, err := p.Products.Search(ctx, keyword)
	if err != nil {
		return nil, swcerr.Service("keyword error", err)
	}

	var productSchema = []dtos.ProductResponse{}
	for _, product := range _products {

		category, err := p.Category.GetByID(ctx, product.CategoryID)
		if err != nil {
			return nil, err
		}

		resp := dtos.ProductResponse{
			ID:          product.ID,
			Price:       product.Price,
			Description: product.Description,
			Name:        product.Name,
			Status:      product.Status,
			Image:       "",
			Created:     utils.HanoiTimezone(product.Created),
			Category:    category.Name,
		}

		if len(strings.Split(product.Image, ",")) > 0 {
			resp.Image = strings.Split(product.Image, ",")[0]
		}

		productSchema = append(productSchema, resp)
	}
	return productSchema, nil
}

// GetProducts implements IProductService.
func (p *Products) GetProducts(ctx context.Context, limit int) ([]dtos.ProductResponse, error) {
	products, err := p.Products.GetLimit(ctx, limit, 1)
	if err != nil {
		return nil, err
	}

	var productResponse = []dtos.ProductResponse{}
	for _, _product := range products {
		var (
			product = dtos.ProductResponse{
				ID:          _product.ID,
				Price:       _product.Price,
				Description: _product.Description,
				Name:        _product.Name,
				Status:      _product.Status,
				Created:     utils.HanoiTimezone(_product.Created),
				Image:       "",
			}
			types enum.Category
		)

		if len(strings.Split(_product.Image, ",")) > 0 {
			product.Image = strings.Split(_product.Image, ",")[0]
		}

		category, err := p.Category.GetByID(ctx, _product.CategoryID)
		if err != nil {
			return nil, err
		}

		if err := types.Load(category.Name); err != nil {
			return nil, fmt.Errorf("[code %d] %v", http.StatusBadRequest, err)
		}

		product.Category = category.Name
		productResponse = append(productResponse, product)
	}
	return productResponse, nil
}
