package controller

import (
	"net/http"
	"strconv"
	"swclabs/swix/app"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/service/paydeliver"

	"github.com/labstack/echo/v4"
)

var _ = app.Controller(NewPaydeliver)

// NewPaydeliver creates a new Products object
func NewPaydeliver(services paydeliver.IPaymentDelivery) IPaydeliver {
	return &Paydeliver{
		Services: services,
	}
}

// IPaydeliver interface for paydeliver controller
type IPaydeliver interface {
	CreateDeliveryAddress(e echo.Context) error
	GetDeliveryAddress(e echo.Context) error
	CreateDelivery(e echo.Context) error
	GetDelivery(e echo.Context) error
}

// Paydeliver struct implementation of IProducts
type Paydeliver struct {
	Services paydeliver.IPaymentDelivery
}

// CreateDelivery .
// @Description create delivery info.
// @Tags delivery
// @Accept json
// @Produce json
// @Param addr body dtos.DeliveryBody true "delivery info request"
// @Success 200 {object} dtos.OK
// @Router /delivery [POST]
func (p *Paydeliver) CreateDelivery(e echo.Context) error {
	var body dtos.DeliveryBody
	if err := e.Bind(&body); err != nil {
		return e.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if err := p.Services.CreateDelivery(e.Request().Context(), body); err != nil {
		return e.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return e.JSON(http.StatusOK, dtos.OK{
		Msg: "your delivery info has been saved",
	})
}

// GetDelivery .
// @Description get delivery info by user id.
// @Tags delivery
// @Accept json
// @Produce json
// @Param uid query string true "user id"
// @Success 200 {object} dtos.OK
// @Router /delivery [GET]
func (p *Paydeliver) GetDelivery(e echo.Context) error {
	uid, err := strconv.Atoi(e.QueryParam("uid"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "missing 'uid' required",
		})
	}

	del, err := p.Services.GetDelivery(e.Request().Context(), int64(uid))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return e.JSON(http.StatusOK, del)
}

// GetDeliveryAddress .
// @Description get address delivery.
// @Tags address
// @Accept json
// @Produce json
// @Param uid query string true "user id"
// @Success 200 {object} []dtos.Address
// @Router /address [GET]
func (p *Paydeliver) GetDeliveryAddress(e echo.Context) error {
	uid, err := strconv.Atoi(e.QueryParam("uid"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, dtos.Error{
			Msg: "missing 'uid' required",
		})
	}

	addr, err := p.Services.GetDeliveryAddress(e.Request().Context(), int64(uid))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return e.JSON(http.StatusOK, addr)
}

// CreateDeliveryAddress .
// @Description create address delivery.
// @Tags address
// @Accept json
// @Produce json
// @Param addr body dtos.DeliveryAddress true "address request"
// @Success 200 {object} []dtos.ProductResponse
// @Router /address [POST]
func (p *Paydeliver) CreateDeliveryAddress(e echo.Context) error {
	var addr dtos.DeliveryAddress
	if err := e.Bind(&addr); err != nil {
		return e.JSON(http.StatusBadRequest, dtos.Error{
			Msg: err.Error(),
		})
	}
	if err := p.Services.CreateDeliveryAddress(e.Request().Context(), addr); err != nil {
		return e.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return e.JSON(http.StatusOK, dtos.OK{
		Msg: "your address has been saved",
	})
}
