package app

import (
	"swclabs/swiftcart/app/middleware"
	"swclabs/swiftcart/app/router"
	"time"

	"swclabs/swiftcart/internal/config"
	"swclabs/swiftcart/internal/service"
	"swclabs/swiftcart/pkg/job"
	"swclabs/swiftcart/pkg/mailers"
	"swclabs/swiftcart/pkg/sentry"

	_ "swclabs/swiftcart/internal/tasks/plugin"
)

func init() {
	sentry.Init()
	mailers.Config(config.Email, config.EmailAppPassword)
}

func (server *Server) scheduler() {
	newJob := job.New()
	go newJob.Scheduler(service.Ping, 5*time.Second)

	newJob.Info()
}

func (server *Server) Run(addr string) error {
	server.scheduler()
	server.backgroundTask()
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
