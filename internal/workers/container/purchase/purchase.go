package purchase

import (
	"swclabs/swipex/app"
	"swclabs/swipex/internal/workers/container/healthcheck"
	"swclabs/swipex/internal/workers/server"
)

func New(base healthcheck.IRouter, router IRouter) app.IApplication {
	mux := server.NewServeMux()
	mux.Handle(base)
	mux.Handle(router)
	return server.New(mux)
}
