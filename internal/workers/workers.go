package workers

import (
	"swclabs/swipex/app"
	"swclabs/swipex/internal/workers/container/authentication"
	"swclabs/swipex/internal/workers/container/healthcheck"
	"swclabs/swipex/internal/workers/container/purchase"
	"swclabs/swipex/internal/workers/server"
)

func NewApp(
	base healthcheck.IRouter,
	auth authentication.IRouter,
	purchase purchase.IRouter,
) app.IApplication {
	mux := server.NewServeMux()
	mux.Handle(base)
	mux.Handle(auth)
	mux.Handle(purchase)
	worker := server.New(mux)
	return worker
}
