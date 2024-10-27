package healthcheck

import (
	"swclabs/swipex/pkg/components"

	"github.com/labstack/echo/v4"
)

// Home is the home page controller
func Home(c echo.Context) error {
	return components.HomeIndex().Render(c.Request().Context(), c.Response())
}
