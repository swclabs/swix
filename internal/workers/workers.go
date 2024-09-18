package workers

import (
	"swclabs/swix/app"
	"swclabs/swix/internal/workers/router"
	"swclabs/swix/internal/workers/server"
)

func NewWorkerNode(
	base router.IBase,
	manager router.IManager,
	purchase router.IPurchase,
) app.IApplication {
	mux := server.NewServeMux()
	mux.Handle(base)
	mux.Handle(manager)
	mux.Handle(purchase)
	worker := server.New(mux)
	return worker
}
