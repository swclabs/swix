// Package purchase implements the purchase interface
package purchase

import (
	"context"
	"fmt"
	"log"
	"strings"
	"swclabs/swix/app"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/domain/xdto"
	"swclabs/swix/internal/core/repos/addresses"
	"swclabs/swix/internal/core/repos/carts"
	"swclabs/swix/internal/core/repos/categories"
	"swclabs/swix/internal/core/repos/deliveries"
	"swclabs/swix/internal/core/repos/inventories"
	"swclabs/swix/internal/core/repos/orders"
	"swclabs/swix/internal/core/repos/products"
	"swclabs/swix/internal/core/repos/specifications"
	"swclabs/swix/internal/core/repos/users"
	"swclabs/swix/internal/core/x/ghnx"
	"swclabs/swix/pkg/infra/db"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// New creates a new Purchase object
var New = app.Service(
	func(
		order orders.IOrders,
		cart carts.ICarts,
		user users.IUsers,
		spec specifications.ISpecifications,
		inv inventories.IInventories,
		product products.IProducts,
		category categories.ICategories,
		address addresses.IAddress,
		delivery deliveries.IDeliveries,
		ghn ghnx.IGhnx,
	) IPurchase {
		return &Purchase{
			Cart:      cart,
			Order:     order,
			User:      user,
			Spec:      spec,
			Inventory: inv,
			Product:   product,
			Category:  category,
			Address:   address,
			Delivery:  delivery,
			Ghn:       ghn,
		}
	},
)

// Purchase struct for purchase service
type Purchase struct {
	Order     orders.IOrders
	Cart      carts.ICarts
	User      users.IUsers
	Spec      specifications.ISpecifications
	Category  categories.ICategories
	Product   products.IProducts
	Inventory inventories.IInventories
	Address   addresses.IAddress
	Delivery  deliveries.IDeliveries
	Ghn       ghnx.IGhnx
}

// AddressDistrict implements IPurchase.
func (p *Purchase) AddressDistrict(ctx context.Context, provinceID int) (*xdto.DistrictDTO, error) {
	return p.Ghn.District(ctx, provinceID)
}

// AddressProvince implements IPurchase.
func (p *Purchase) AddressProvince(ctx context.Context) (*xdto.ProvinceDTO, error) {
	return p.Ghn.Province(ctx)
}

// AddressWard implements IPurchase.
func (p *Purchase) AddressWard(ctx context.Context, districtID int) (*xdto.WardDTO, error) {
	return p.Ghn.Ward(ctx, districtID)
}

// CreateDelivery implements IPurchase.
func (p *Purchase) CreateDelivery(ctx context.Context, delivery dtos.DeliveryBody) error {
	sendate, err := time.Parse(time.RFC3339, delivery.SentDate)
	if err != nil {
		sendate = time.Time{}
	}
	receivedate, err := time.Parse(time.RFC3339, delivery.ReceivedDate)
	if err != nil {
		receivedate = time.Time{}
	}
	return p.Delivery.Create(ctx, entity.Deliveries{
		UserID:       delivery.UserID,
		AddressID:    delivery.AddressID,
		Status:       delivery.Status,
		Method:       delivery.Method,
		Note:         delivery.Note,
		SentDate:     sendate,
		ReceivedDate: receivedate,
	})
}

// CreateDeliveryAddress implements IPurchase.
func (p *Purchase) CreateDeliveryAddress(ctx context.Context, addr dtos.DeliveryAddress) error {
	return p.Address.Insert(ctx, entity.Addresses{
		UserID:   addr.UserID,
		Street:   addr.Street,
		City:     addr.City,
		Ward:     addr.Ward,
		District: addr.District,
	})
}

// GetDelivery implements IPurchase.
func (p *Purchase) GetDelivery(ctx context.Context, userID int64) ([]dtos.Delivery, error) {
	deliveries, err := p.Delivery.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var delivery = []dtos.Delivery{}
	for _, del := range deliveries {
		var (
			sentdate     string
			receiveddate string
		)
		if !del.SentDate.IsZero() {
			sentdate = del.SentDate.Format(time.RFC3339)
		}
		if !del.ReceivedDate.IsZero() {
			receiveddate = del.ReceivedDate.Format(time.RFC3339)
		}
		delivery = append(delivery, dtos.Delivery{
			ID:           del.ID,
			AddressID:    del.AddressID,
			UserID:       del.UserID,
			Status:       del.Status,
			Method:       del.Method,
			Note:         del.Note,
			SentDate:     sentdate,
			ReceivedDate: receiveddate,
		})
	}
	return delivery, nil
}

// GetDeliveryAddress implements IPurchase.
func (p *Purchase) GetDeliveryAddress(ctx context.Context, userID int64) ([]dtos.Address, error) {
	addrs, err := p.Address.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var addresses = []dtos.Address{}
	for _, addr := range addrs {
		addresses = append(addresses, dtos.Address{
			ID:       addr.ID,
			Street:   addr.Street,
			City:     addr.City,
			Ward:     addr.Ward,
			District: addr.District,
		})
	}
	return addresses, nil
}

// GetOrdersByUserID implements IPurchaseService.
func (p *Purchase) GetOrdersByUserID(ctx context.Context, userID int64, limit int) ([]dtos.OrderSchema, error) {
	// Get orders by user ID
	orders, err := p.Order.Get(ctx, userID, limit)
	if err != nil {
		return nil, err
	}
	// Get user by user ID
	user, err := p.User.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	var orderSchema = []dtos.OrderSchema{}

	for _, order := range orders {
		// Get products by order ID
		products, err := p.Order.GetProductByOrderID(ctx, order.ID)
		if err != nil {
			return nil, err
		}
		var productSchema []dtos.ProductOrderSchema
		for _, product := range products {
			productSchema = append(productSchema, dtos.ProductOrderSchema{
				ID:           product.ID,
				OrderID:      product.OrderID,
				CurrencyCode: product.CurrencyCode,
				InventoryID:  product.InventoryID,
				Quantity:     product.Quantity,
				TotalAmount:  product.TotalAmount.String(),
			})
		}
		// Merge product and order schema
		orderSchema = append(orderSchema, dtos.OrderSchema{
			ID:        order.ID,
			UserID:    user.ID,
			Time:      order.Time.Format(time.RFC3339),
			UUID:      order.UUID,
			Status:    order.Status,
			Items:     productSchema,
			UserEmail: user.Email,
			Username:  fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		})
	}
	return orderSchema, nil
}

// DeleteItemFromCart implements IPurchaseService.
func (p *Purchase) DeleteItemFromCart(ctx context.Context, cartID int64) error {
	return p.Cart.RemoveItem(ctx, cartID)
}

// AddToCart implements IPurchaseService.
func (p *Purchase) AddToCart(ctx context.Context, cart dtos.CartInsert) error {
	specs, err := p.Spec.GetByID(ctx, cart.SpecID)
	if err != nil {
		return err
	}
	return p.Cart.Insert(ctx, entity.Carts{
		UserID:      cart.UserID,
		InventoryID: cart.InventoryID,
		Quantity:    cart.Quantity,
		SpecID:      specs.ID,
	})
}

// GetCart implements IPurchaseService.
func (p *Purchase) GetCart(ctx context.Context, userID int64, limit int) (*dtos.CartSlices, error) {
	carts, err := p.Cart.GetCartByUserID(ctx, userID, limit)
	if err != nil {
		return nil, err
	}
	var cartSchema = dtos.CartSlices{
		UserID: userID,
	}
	for _, item := range carts {
		inv, err := p.Inventory.GetByID(ctx, item.InventoryID)
		if err != nil {
			return nil, err
		}
		prod, err := p.Product.GetByID(ctx, inv.ProductID)
		if err != nil {
			return nil, err
		}
		category, err := p.Category.GetByID(ctx, prod.CategoryID)
		if err != nil {
			return nil, err
		}
		var (
			amount = decimal.NewFromUint64(uint64(item.Quantity)).Mul(inv.Price)
			images = strings.Split(inv.Image, ",")
			image  = ""
		)
		if len(images) != 0 {
			image = images[0]
		}
		cartSchema.Products = append(cartSchema.Products, dtos.CartSchema{
			ID:          item.ID,
			Quantity:    item.Quantity,
			ProductName: prod.Name,
			Amount:      amount.String(),
			Img:         image,
			Category:    category.Name,
		})
	}
	return &cartSchema, nil
}

// CreateOrders implements IPurchaseService.
func (p *Purchase) CreateOrders(ctx context.Context, createOrder dtos.CreateOrderSchema) (string, error) {
	tx, err := db.NewTransaction(ctx)
	if err != nil {
		return "", err
	}
	var (
		_uuid              = strings.ReplaceAll(uuid.New().String(), "_", " ")
		inventoryRepo      = inventories.New(tx)
		orderRepo          = orders.New(tx)
		totalAmount        decimal.Decimal
		productTotalAmount []decimal.Decimal
	)

	for _, product := range createOrder.Products {
		inven, err := inventoryRepo.GetByID(ctx, product.InventoryID)
		if err != nil {
			if errTx := tx.Rollback(ctx); errTx != nil {
				log.Fatal(errTx)
			}
			return "", err
		}
		totalAmount = totalAmount.Add(inven.Price)
		productTotalAmount = append(
			productTotalAmount,
			inven.Price.Mul(decimal.NewFromInt32(int32(product.Quantity))))
	}
	orderID, err := orderRepo.Create(ctx, entity.Orders{
		UUID:        _uuid,
		DeliveryID:  createOrder.DeleveryID,
		UserID:      createOrder.UserID,
		Status:      "pending",
		TotalAmount: totalAmount,
	})
	if err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return "", err
	}
	for idx, product := range createOrder.Products {
		if err := orderRepo.InsertProduct(ctx, entity.ProductInOrder{
			OrderID:     orderID,
			InventoryID: product.InventoryID,
			Quantity:    product.Quantity,
			TotalAmount: productTotalAmount[idx],
		}); err != nil {
			if errTx := tx.Rollback(ctx); errTx != nil {
				log.Fatal(errTx)
			}
			return "", err
		}
	}
	return _uuid, tx.Commit(ctx)
}
