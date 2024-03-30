package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"swclabs/swipe-api/pkg/web/components"
)

func main() {
	app := echo.New()
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	app.GET("/purchase-order", func(c echo.Context) error {
		return components.PurchaseOrderIndex().Render(c.Request().Context(), c.Response())
	})
	app.Logger.Fatal(app.Start("localhost:8000"))
}
