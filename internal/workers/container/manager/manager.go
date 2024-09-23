package manager

import (
	"swclabs/swix/app"
	"swclabs/swix/internal/workers/container/base"
	"swclabs/swix/internal/workers/server"
)

func New(base base.IRouter, router IRouter) app.IApplication {
	mux := server.NewServeMux()
	mux.Handle(base)
	mux.Handle(router)
	return server.New(mux)
}
