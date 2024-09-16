package workers

import (
	"swclabs/swix/boot"
	"swclabs/swix/internal/workers/router"
	"swclabs/swix/internal/workers/server"
)

func NewWorkerNode(
	base router.IBase,
	manager router.IManager,
	purchase router.IPurchase,
) boot.ICore {
	mux := server.NewServeMux()
	mux.Handle(base)
	mux.Handle(manager)
	mux.Handle(purchase)
	worker := server.New(mux)
	return worker
}
