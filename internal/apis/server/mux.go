package server

import (
	"swclabs/swix/internal/apis/router"

	"github.com/labstack/echo/v4"
)

type IMux interface {
	Handle(router router.IRouter)
	ServeHTTP(engine *echo.Echo)
}

func NewServeMux() IMux {
	return &Mux{}
}

type Mux struct {
	router []router.IRouter
}

// Handle implements IMux.
func (m *Mux) Handle(router router.IRouter) {
	m.router = append(m.router, router)
}

// ServeHTTP implements IMux.
func (m *Mux) ServeHTTP(engine *echo.Echo) {
	for _, r := range m.router {
		r.Routers(engine)
	}
}
