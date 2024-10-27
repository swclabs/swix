package article

import (
	"swclabs/swipex/app"
	"swclabs/swipex/internal/apis/container/healthcheck"
	"swclabs/swipex/internal/apis/server"
)

func New(base healthcheck.IRouter, router IRouter) app.IApplication {
	mux := server.NewServeMux()
	mux.Handle(base)
	mux.Handle(router)
	return server.New(mux)
}
