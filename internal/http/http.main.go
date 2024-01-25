package http

import (
	"github.com/swclabs/swipe-api/internal/http/middleware"

	"github.com/swclabs/swipe-api/internal/config"
	"github.com/swclabs/swipe-api/pkg/mailers"
	"github.com/swclabs/swipe-api/pkg/sentry"
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
}

func (server *_Server) Bootstrap(fn ...func(server IServer)) {
	for _, _fn := range fn {
		_fn(server)
	}
}

func (server *_Server) Run(addr string) error {
	return server.engine.Run(addr)
}
