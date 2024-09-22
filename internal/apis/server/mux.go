package server

import (
	"github.com/labstack/echo/v4"
)

type IMux interface {
	Handle(router IRouter)
	ServeHTTP(engine *echo.Echo)
}

type IRouter interface {
	Routers(e *echo.Echo)
}

func NewServeMux() IMux {
	return &Mux{}
}

type Mux struct {
	router []IRouter
}

// Handle implements IMux.
func (m *Mux) Handle(router IRouter) {
	m.router = append(m.router, router)
}

// ServeHTTP implements IMux.
func (m *Mux) ServeHTTP(engine *echo.Echo) {
	for _, r := range m.router {
		r.Routers(engine)
	}
}
