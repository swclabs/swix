package http

import (
	"swclabs/swiftcart/internal/http/middleware"
	"swclabs/swiftcart/internal/http/router"
	"swclabs/swiftcart/internal/misc/resolver"

	"swclabs/swiftcart/internal/config"
	"swclabs/swiftcart/pkg/mailers"
	"swclabs/swiftcart/pkg/sentry"
)

var _ IServer = &_Server{}

func init() {
	sentry.Init()
	mailers.Config(config.Email, config.EmailAppPassword)
}

func (server *_Server) prepare() {
	server.middleware(
		middleware.GinMiddleware,
		middleware.Sentry,
	)
	server.router(
		router.Common,
		router.Docs,
	)
}

func (server *_Server) InitAccountManagement() {
	server.backgroundTask(func() {
		resolver.StartImageHandler(5)
	})
	var accountManagement = router.NewAccountManagement()
	server.router(
		accountManagement.Users,
		accountManagement.Auth,
		accountManagement.OAuth2,
	)
}

func (server *_Server) Run(addr string) error {
	return server.engine.Run(addr)
}
