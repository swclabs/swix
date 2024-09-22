package classify

import (
	"swclabs/swix/app"
	"swclabs/swix/internal/apis/container/base"
	"swclabs/swix/internal/apis/server"
)

func New(base base.IRouter, router IRouter) app.IApplication {
	mux := server.NewServeMux()
	mux.Handle(base)
	mux.Handle(router)
	return server.New(mux)
}
