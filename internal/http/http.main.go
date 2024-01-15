package http

import (
	"swclabs/swiftcart/internal/http/middleware"
	"swclabs/swiftcart/internal/http/router"
	"swclabs/swiftcart/internal/misc/resolver"

	"swclabs/swiftcart/internal/config"
	"swclabs/swiftcart/pkg/mailers"
	"swclabs/swiftcart/pkg/sentry"
)

func init() {
	sentry.Init()
	mailers.Config(config.Email, config.EmailAppPassword)
}

func (server *Server) Run(addr string) error {
	server.backgroundTask(func() {
		resolver.StartImageHandler(5)
	})
	server.middleware(
		middleware.GinMiddleware,
		middleware.Sentry,
	)
	server.router(
		router.Common,
		router.Auth,
		router.OAuth2,
		router.Users,
		router.Docs,
	)
	return server.engine.Run(addr)
}
