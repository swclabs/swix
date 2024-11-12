package workers

import (
	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/workers/container/authentication"
	"github.com/swclabs/swipex/internal/workers/container/healthcheck"
	"github.com/swclabs/swipex/internal/workers/container/purchase"
	"github.com/swclabs/swipex/internal/workers/server"
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
