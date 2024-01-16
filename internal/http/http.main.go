package http

import (
	"swclabs/swiftcart/internal/http/middleware"
	"swclabs/swiftcart/internal/http/router"
	"swclabs/swiftcart/internal/misc/resolver"

	"swclabs/swiftcart/internal/config"
	"swclabs/swiftcart/pkg/mailers"
	"swclabs/swiftcart/pkg/sentry"
)

var _ IServer = &Server{}

func init() {
	sentry.Init()
	mailers.Config(config.Email, config.EmailAppPassword)
}

func (server *Server) InitMiddleware() {
	server.middleware(
		middleware.GinMiddleware,
		middleware.Sentry,
	)
	server.router(
		router.Common,
		router.Docs,
	)
}

func (server *Server) InitAccountManagement() {
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

func (server *Server) Run(addr string) error {
	return server.engine.Run(addr)
}
