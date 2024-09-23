package workers

import (
	"swclabs/swix/app"
	"swclabs/swix/internal/workers/container/base"
	"swclabs/swix/internal/workers/container/manager"
	"swclabs/swix/internal/workers/container/purchase"
	"swclabs/swix/internal/workers/server"
)

func NewApp(
	base base.IRouter,
	manager manager.IRouter,
	purchase purchase.IRouter,
) app.IApplication {
	mux := server.NewServeMux()
	mux.Handle(base)
	mux.Handle(manager)
	mux.Handle(purchase)
	worker := server.New(mux)
	return worker
}
