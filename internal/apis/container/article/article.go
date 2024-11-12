package article

import (
	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/apis/container/healthcheck"
	"github.com/swclabs/swipex/internal/apis/server"
)

func New(base healthcheck.IRouter, router IRouter) app.IApplication {
	mux := server.NewServeMux()
	mux.Handle(base)
	mux.Handle(router)
	return server.New(mux)
}
