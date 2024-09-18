// Package controller implements the controller interface
package controller

import (
	"net/http"
	"swclabs/swix/boot"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/service/base"
	"swclabs/swix/internal/core/x/oauth2"
	"swclabs/swix/pkg/utils"

	"github.com/labstack/echo/v4"
)

var _ = boot.Controller(New)

// New creates a new Base object
func New(services base.IService) IBaseController {
	return &BaseController{
		service: services,
	}
}

// IBaseController interface for base controller
type IBaseController interface {
	HealthCheck(c echo.Context) error
	WorkerCheck(c echo.Context) error
}

// BaseController struct implementation of IBase
type BaseController struct {
	service base.IService
}

// HealthCheck .
// @Description health check api server.
// @Tags base
// @Accept json
// @Produce json
// @Success 200
// @Router /common/healthcheck [GET]
func (b *BaseController) HealthCheck(c echo.Context) error {
	return c.JSON(200, b.service.HealthCheck(c.Request().Context()))
}

// WorkerCheck .
// @Description health check worker consume server.
// @Tags base
// @Accept json
// @Produce json
// @Success 200
// @Router /common/worker [GET]
func (b *BaseController) WorkerCheck(c echo.Context) error {
	results, err := base.UseTask(b.service).WorkerCheckResult(c.Request().Context(), 10)
	if err != nil {
		return c.JSON(400, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(200, dtos.OK{
		Msg: results,
	})
}

// Auth0Login .
// @Description Auth0 Login form.
// @Tags base
// @Accept json
// @Produce json
// @Success 200
// @Router /oauth2/login [GET]
func Auth0Login(c echo.Context) error {
	auth := oauth2.New()
	url := auth.AuthCodeURL(auth.State)
	if err := utils.SaveSession(c, utils.BaseSessions, "state", auth.State); err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.Error{
			Msg: err.Error(),
		})
	}
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

// Auth0Callback .
func Auth0Callback(c echo.Context) error {
	auth := oauth2.New()
	return auth.OAuth2CallBack(c)
}
