package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/service/purchase"
	"swclabs/swipecore/pkg/lib/valid"

	"github.com/labstack/echo/v4"
)

type IPurchase interface {
	AddToCarts(c echo.Context) error
	GetCarts(c echo.Context) error
	DeleteItem(c echo.Context) error
}

type Purchase struct {
	services purchase.IPurchaseService
}

var _ IPurchase = (*Purchase)(nil)

func NewPurchase(services purchase.IPurchaseService) IPurchase {
	return &Purchase{services: services}
}

// AddToCarts
// @Description add item to carts.
// @Tags purchase
// @Accept json
// @Produce json
// @Param login body domain.CartInsertReq true "cart insert request"
// @Success 200 {object} domain.OK
// @Router /purchase/carts [POST]
func (purchase *Purchase) AddToCarts(c echo.Context) error {
	var cartReq domain.CartInsertReq
	if err := c.Bind(&cartReq); err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err.Error(),
		})
	}
	if err := valid.Validate(&cartReq); err != "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: err,
		})
	}
	if err := purchase.services.AddToCart(c.Request().Context(), cartReq); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, domain.OK{
		Msg: "your item has been update successfully",
	})
}

// GetCarts
// @Description get list of items from carts
// @Tags purchase
// @Accept json
// @Produce json
// @Param uid query string true "user id"
// @Success 200 {object} domain.CartSchema
// @Router /purchase/carts [GET]
func (purchase *Purchase) GetCarts(c echo.Context) error {
	sUserId := c.QueryParam("uid")
	if sUserId == "" {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "missing 'uid' required",
		})
	}
	userId, err := strconv.ParseInt(sUserId, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{
			Msg: "'uid' must be a positive integer",
		})
	}
	items, err := purchase.services.GetCart(c.Request().Context(), userId, 10)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, *items)
}

// DeleteItem
// @Description delete item from carts
// @Tags purchase
// @Accept json
// @Produce json
// @Param uid query int true "user id"
// @Param wid query int true "inventory id"
// @Success 200 {object} domain.OK
// @Router /purchase/carts [DELETE]
func (purchase *Purchase) DeleteItem(c echo.Context) error {
	var (
		ids  = make(map[string]string)
		iIds = make(map[string]int64)
	)
	const (
		uid = "uid"
		wid = "wid"
	)
	ids[uid] = c.QueryParam(uid)
	ids[wid] = c.QueryParam(wid)

	for key, param := range ids {
		if param == "" {
			return c.JSON(http.StatusBadRequest, domain.Error{
				Msg: fmt.Sprintf("missing param '%s' required", key),
			})
		}
		id, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, domain.Error{
				Msg: fmt.Sprintf(" param '%s' must be integer", key),
			})
		}
		iIds[key] = id
	}

	if err := purchase.services.
		DeleteItemFromCart(c.Request().Context(), iIds[uid], iIds[wid]); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, domain.OK{
		Msg: "your item has been deleted successfully",
	})
}
