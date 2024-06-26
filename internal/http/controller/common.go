package controller

import (
	"net/http"
	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/internal/core/service/common"
	"swclabs/swipecore/internal/core/utils/oauth2"
	"swclabs/swipecore/pkg/utils"

	"github.com/labstack/echo/v4"
)

type ICommon interface {
	HealthCheck(c echo.Context) error
	WorkerCheck(c echo.Context) error
}

type Common struct {
	service common.ICommonService
}

func NewCommon(services common.ICommonService) ICommon {
	return &Common{
		service: services,
	}
}

// HealthCheck .
// @Description health check api server.
// @Tags common
// @Accept json
// @Produce json
// @Success 200
// @Router /common/healthcheck [GET]
func (common *Common) HealthCheck(c echo.Context) error {
	return c.JSON(200, common.service.HealthCheck(c.Request().Context()))
}

// Auth0Login .
// @Description Auth0 Login form.
// @Tags common
// @Accept json
// @Produce json
// @Success 200
// @Router /oauth2/login [GET]
func Auth0Login(c echo.Context) error {
	auth := oauth2.New()
	url := auth.AuthCodeURL(auth.State)
	if err := utils.SaveSession(c, utils.BaseSessions, "state", auth.State); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func Auth0Callback(c echo.Context) error {
	auth := oauth2.New()
	return auth.OAuth2CallBack(c)
}

// WorkerCheck .
// @Description health check worker consume server.
// @Tags common
// @Accept json
// @Produce json
// @Success 200
// @Router /common/worker [GET]
func (common *Common) WorkerCheck(c echo.Context) error {
	results, err := common.service.CallTask().WorkerCheckResult(c.Request().Context(), 10)
	if err != nil {
		return c.JSON(400, domain.Error{
			Msg: err.Error(),
		})
	}
	return c.JSON(200, domain.OK{
		Msg: results,
	})
}

func Foo(ctx echo.Context) error {
	// sentrygin handler will catch it just fine. Also, because we attached "someRandomTag"
	// in the middleware before, it will be sent through as well
	panic("y tho")
}
