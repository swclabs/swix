package http

import (
	"github.com/swclabs/swipe-api/internal/http/middleware"

	"github.com/swclabs/swipe-api/internal/config"
	"github.com/swclabs/swipe-api/pkg/sentry"
	"github.com/swclabs/swipe-api/pkg/tools/mailers"
)

var _ IServer = &_Server{}

func init() {
	sentry.Init()
	mailers.Config(config.Email, config.EmailAppPassword)
}

func (server *_Server) initMiddleware() {
	server.middleware(
		middleware.BaseSetting,
		middleware.CookieSetting,
		middleware.Sentry,
	)
}

func (server *_Server) Bootstrap(fn ...func(server IServer)) {
	for _, _fn := range fn {
		_fn(server)
	}
}

func (server *_Server) Run(addr string) error {
	return server.engine.Start(addr)
}
