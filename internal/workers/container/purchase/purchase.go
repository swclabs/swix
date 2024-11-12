package purchase

import (
	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/workers/container/healthcheck"
	"github.com/swclabs/swipex/internal/workers/server"
)

func New(base healthcheck.IRouter, router IRouter) app.IApplication {
	mux := server.NewServeMux()
	mux.Handle(base)
	mux.Handle(router)
	return server.New(mux)
}
