// Package purchase implements the purchase interface
package purchase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"swclabs/swipex/app"
	"swclabs/swipex/internal/core/domain/dtos"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/internal/core/domain/xdto"
	"swclabs/swipex/internal/core/repos/addresses"
	"swclabs/swipex/internal/core/repos/carts"
	"swclabs/swipex/internal/core/repos/categories"
	"swclabs/swipex/internal/core/repos/coupons"
	"swclabs/swipex/internal/core/repos/deliveries"
	"swclabs/swipex/internal/core/repos/inventories"
	"swclabs/swipex/internal/core/repos/orders"
	"swclabs/swipex/internal/core/repos/products"
	"swclabs/swipex/internal/core/repos/users"
	"swclabs/swipex/internal/core/x/ghnx"
	"swclabs/swipex/pkg/infra/db"
	"swclabs/swipex/pkg/utils"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/shopspring/decimal"
)

// New creates a new Purchase object
var New = app.Service(
	func(
		order orders.IOrders,
		cart carts.ICarts,
		user users.IUsers,
		inv inventories.IInventories,
		product products.IProducts,
		category categories.ICategories,
		address addresses.IAddress,
		delivery deliveries.IDeliveries,
		ghn ghnx.IGhnx,
		coupon coupons.ICoupons,
	) IPurchase {
		return &Purchase{
			Coupon:    coupon,
			Cart:      cart,
			Order:     order,
			User:      user,
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
	Coupon    coupons.ICoupons
	Order     orders.IOrders
	Cart      carts.ICarts
	User      users.IUsers
	Category  categories.ICategories
	Product   products.IProducts
	Inventory inventories.IInventories
	Address   addresses.IAddress
	Delivery  deliveries.IDeliveries
	Ghn       ghnx.IGhnx
}

// CreateCoupon implements IPurchase.
func (p *Purchase) CreateCoupon(ctx context.Context, coupon dtos.CreateCoupon) (code string, err error) {
	code = utils.GenCouponsCode(10)
	exp, err := time.Parse(time.RFC3339, coupon.ExpiredAt)
	if err != nil {
		return "", err
	}
	err = p.Coupon.Create(ctx, entity.Coupons{
		Code:        code,
		Discount:    coupon.Discount,
		Status:      coupon.Status,
		Used:        0,
		MaxUse:      coupon.MaxUse,
		Description: coupon.Description,
		ExpiredAt:   exp,
	})
	return
}

// GetCoupon implements IPurchase.
func (p *Purchase) GetCoupon(ctx context.Context) (coupons []dtos.Coupon, err error) {
	_coupons, err := p.Coupon.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	for _, coupon := range _coupons {
		if coupon.Status == "active" {
			coupons = append(coupons, dtos.Coupon{
				Code:        coupon.Code,
				Discount:    coupon.Discount,
				ID:          coupon.ID,
				ExpiredAt:   coupon.ExpiredAt,
				Description: coupon.Description,
			})
		}
	}
	return coupons, nil
}

// UseCoupon implements IPurchase.
func (p *Purchase) UseCoupon(ctx context.Context, userID int64, couponCode string) error {
	coupons, err := p.Coupon.GetByUser(ctx, userID)
	if !errors.Is(err, pgx.ErrNoRows) {
		return err
	}
	if len(coupons) > 0 {
		for _, coupon := range coupons {
			if coupon.CouponCode == couponCode {
				return errors.New("coupon already used")
			}
		}
	}
	return nil
}

// GetOrderByCode implements IPurchase.
func (p *Purchase) GetOrderByCode(ctx context.Context, orderCode string) (*dtos.OrderInfo, error) {
	items, err := p.Order.GetItemByCode(ctx, orderCode)
	if err != nil {
		return nil, fmt.Errorf("error getting order by UUID: %w", err)
	}
	order, err := p.Order.GetByUUID(ctx, orderCode)
	if err != nil {
		return nil, err
	}
	if len(items) >= 1 {
		user, err := p.User.GetByID(ctx, order.UserID)
		if err != nil {
			return nil, err
		}
		delivery, err := p.Delivery.GetByID(ctx, order.DeliveryID)
		if err != nil {
			return nil, err
		}
		address, err := p.Address.GetByID(ctx, delivery.AddressID)
		if err != nil {
			return nil, err
		}
		return &dtos.OrderInfo{
			Items:     items,
			UUID:      order.UUID,
			CreatedAt: utils.HanoiTimezone(order.Time),
			User: dtos.OrderFormCustomer{
				Email:     user.Email,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Phone:     user.PhoneNumber,
			},
			Delivery: dtos.OrderFormDelivery{
				Status:   order.Status,
				Method:   delivery.Method,
				Note:     delivery.Note,
				SentDate: utils.HanoiTimezone(delivery.SentDate),
			},
			Address: dtos.OrderFormAddress{
				City:     address.City,
				Ward:     address.Ward,
				District: address.District,
				Street:   address.Street,
			},
		}, nil
	}
	return nil, errors.New("order not found")
}

// CreateOrderForm implements IPurchase.
func (p *Purchase) CreateOrderForm(ctx context.Context, order dtos.OrderForm) (string, error) {
	tx, err := db.NewTx(ctx)
	if err != nil {
		return "", err
	}
	var (
		userRepo      = users.New(tx)
		addressRepo   = addresses.New(tx)
		orderRepo     = orders.New(tx)
		deliveryRepo  = deliveries.New(tx)
		inventoryRepo = inventories.New(tx)
	)
	var (
		totalAmount        decimal.Decimal
		productTotalAmount []decimal.Decimal
		_uuid              = utils.GenOrderCode(16)
	)
	user, err := userRepo.GetByEmail(ctx, order.Customer.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			user = &entity.User{
				Email:       order.Customer.Email,
				FirstName:   order.Customer.FirstName,
				LastName:    order.Customer.LastName,
				PhoneNumber: order.Customer.Phone,
			}
			if user.ID, err = userRepo.Insert(ctx, *user); err != nil {
				if errTx := tx.Rollback(ctx); errTx != nil {
					log.Fatal(errTx)
				}
			}
		} else {
			if errTx := tx.Rollback(ctx); errTx != nil {
				log.Fatal(errTx)
			}
		}
	}
	addrID, err := addressRepo.Insert(ctx, entity.Address{
		UserID:   user.ID,
		Street:   order.Address.Street,
		City:     order.Address.City,
		Ward:     order.Address.Ward,
		District: order.Address.District,
	})
	if err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
	}
	deliveryID, err := deliveryRepo.Create(ctx, entity.Delivery{
		UserID:    user.ID,
		AddressID: addrID,
		Status:    order.Delivery.Status,
		Method:    order.Delivery.Method,
		Note:      order.Delivery.Note,
		SentDate:  time.Now().UTC(),
	})
	if err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
	}
	for {
		_order, err := orderRepo.GetByUUID(ctx, _uuid)
		if err != nil || _order.UUID == "" {
			break
		}
		_uuid = utils.GenOrderCode(16)
	}
	for _, product := range order.Product {
		code := strings.Split(product.Code, "#")
		if len(code) != 2 {
			if errTx := tx.Rollback(ctx); errTx != nil {
				log.Fatal(errTx)
			}
			return "", fmt.Errorf("invalid product code : %s", product.Code)
		}
		id, err := strconv.ParseInt(code[1], 10, 64)
		if err != nil {
			if errTx := tx.Rollback(ctx); errTx != nil {
				log.Fatal(errTx)
			}
			return "", err
		}
		inven, err := inventoryRepo.GetByID(ctx, id)
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
	orderID, err := orderRepo.Create(ctx, entity.Order{
		UUID:        _uuid,
		DeliveryID:  deliveryID,
		UserID:      user.ID,
		Status:      "pending",
		TotalAmount: totalAmount,
	})
	if err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return "", err
	}
	for idx, product := range order.Product {
		code := strings.Split(product.Code, "#")
		if len(code) != 2 {
			if errTx := tx.Rollback(ctx); errTx != nil {
				log.Fatal(errTx)
			}
			return "", fmt.Errorf("invalid product code: %s", product.Code)
		}
		id, err := strconv.ParseInt(code[1], 10, 64)
		if err != nil {
			if errTx := tx.Rollback(ctx); errTx != nil {
				log.Fatal(errTx)
			}
			return "", err
		}
		if err := orderRepo.InsertProduct(ctx, entity.ProductInOrder{
			OrderID:     orderID,
			InventoryID: id,
			Quantity:    product.Quantity,
			TotalAmount: productTotalAmount[idx],
		}); err != nil {
			if errTx := tx.Rollback(ctx); errTx != nil {
				log.Fatal(errTx)
			}
			return "", err
		}
	}
	return "", tx.Commit(ctx)
}

// CreateDeliveryOrder implements IPurchase.
func (p *Purchase) CreateDeliveryOrder(ctx context.Context, shopID int, order xdto.CreateOrderDTO) (*xdto.OrderDTO, error) {
	return p.Ghn.CreateOrder(ctx, shopID, order)
}

// DeliveryOrderInfo implements IPurchase.
func (p *Purchase) DeliveryOrderInfo(ctx context.Context, orderCode string) (*xdto.OrderInfoDTO, error) {
	return p.Ghn.OrderInfo(ctx, orderCode)
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
	_, err = p.Delivery.Create(ctx, entity.Delivery{
		UserID:    delivery.UserID,
		AddressID: delivery.AddressID,
		Status:    delivery.Status,
		Method:    delivery.Method,
		Note:      delivery.Note,
		SentDate:  sendate,
	})
	return err
}

// CreateDeliveryAddress implements IPurchase.
func (p *Purchase) CreateDeliveryAddress(ctx context.Context, addr dtos.DeliveryAddress) error {
	_, err := p.Address.Insert(ctx, entity.Address{
		UserID:   addr.UserID,
		Street:   addr.Street,
		City:     addr.City,
		Ward:     addr.Ward,
		District: addr.District,
	})
	return err
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
		address, err := p.Address.GetByID(ctx, del.AddressID)
		if err != nil {
			return nil, err
		}
		delivery = append(delivery, dtos.Delivery{
			ID: del.ID,
			Address: dtos.Address{
				ID:       address.ID,
				Street:   address.Street,
				City:     address.City,
				Ward:     address.Ward,
				District: address.District,
			},
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
func (p *Purchase) GetOrdersByUserID(ctx context.Context, userID int64, limit int) ([]dtos.OrderInfo, error) {
	// Get orders by user ID
	orders, err := p.Order.Get(ctx, userID, limit)
	if err != nil {
		return nil, err
	}
	var orderInfo = []dtos.OrderInfo{}
	for _, order := range orders {
		// Get products by order code
		order, err := p.GetOrderByCode(ctx, order.UUID)
		if err != nil {
			return nil, fmt.Errorf("error getting order by UUID: %w", err)
		}
		// Merge product and order schema
		orderInfo = append(orderInfo, *order)
	}
	return orderInfo, nil
}

// DeleteItemFromCart implements IPurchaseService.
func (p *Purchase) DeleteItemFromCart(ctx context.Context, cartID int64) error {
	return p.Cart.RemoveByID(ctx, cartID)
}

// AddToCart implements IPurchaseService.
func (p *Purchase) AddToCart(ctx context.Context, cart dtos.CartInsertDTO) error {
	user, err := p.User.GetByEmail(ctx, cart.Email)
	if err != nil {
		return fmt.Errorf("error getting user by email: %v", err)
	}
	return p.Cart.Insert(ctx, entity.Cart{
		UserID:      user.ID,
		InventoryID: cart.InventoryID,
		Quantity:    cart.Quantity,
	})
}

// GetCart implements IPurchaseService.
func (p *Purchase) GetCart(ctx context.Context, userID int64, limit int) (*dtos.Carts, error) {
	carts, err := p.Cart.GetCartInfo(ctx, userID)
	if err != nil {
		return nil, err
	}
	var cartResp = dtos.Carts{
		UserID:   userID,
		Products: []dtos.Cart{},
	}
	for _, cart := range carts {
		var specs dtos.Specs
		if err := json.Unmarshal([]byte(cart.InventorySpecs), &specs); err != nil {
			return nil, fmt.Errorf("error unmarshal inventory specs: %v", err)
		}
		cartResp.Products = append(cartResp.Products, dtos.Cart{
			Name:           cart.Name,
			CartID:         cart.CartID,
			InventoryID:    cart.InventoryID,
			ProductID:      cart.ProductID,
			Quantity:       cart.Quantity,
			Color:          cart.Color,
			InventoryPrice: cart.InventoryPrice,
			CurrencyCode:   cart.CurrencyCode,
			InventoryImage: cart.InventoryImage,
			CategoryName:   cart.CategoryName,
			InventorySpecs: specs,
			Code:           fmt.Sprintf("%s#%d", cart.CategoryName, cart.InventoryID),
		})
	}
	return &cartResp, nil
}

// CreateOrders implements IPurchaseService.
func (p *Purchase) CreateOrders(ctx context.Context, userID int64, createOrder dtos.Order) (string, error) {
	tx, err := db.NewTx(ctx)
	if err != nil {
		return "", err
	}
	var cartRepo = carts.New(tx)
	code, err := p.CreateOrderForm(ctx, dtos.OrderForm(createOrder))
	if err != nil {
		if errTx := tx.Rollback(ctx); errTx != nil {
			log.Fatal(errTx)
		}
		return "", err
	}
	for _, item := range createOrder.Product {
		itemID := strings.Split(item.Code, "#")
		if len(itemID) != 2 {
			if errTx := tx.Rollback(ctx); errTx != nil {
				log.Fatal(errTx)
			}
			return "", fmt.Errorf("invalid product code: %s", item.Code)
		}
		id, err := strconv.ParseInt(itemID[1], 10, 64)
		if err != nil {
			if errTx := tx.Rollback(ctx); errTx != nil {
				log.Fatal(errTx)
			}
			return "", err
		}
		if err := cartRepo.RemoveByItemID(ctx, userID, id); err != nil {
			if errTx := tx.Rollback(ctx); errTx != nil {
				log.Fatal(errTx)
			}
			return "", err
		}
	}
	return code, tx.Commit(ctx)
}
