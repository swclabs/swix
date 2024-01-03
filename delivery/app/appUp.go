package app

import (
	"time"

	"swclabs/swiftcart/delivery/app/middleware"
	"swclabs/swiftcart/delivery/app/router"
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

func (app *App) scheduler() {
	newJob := job.New()
	go newJob.Scheduler(service.Ping, 5*time.Second)

	newJob.Info()
}

func (app *App) Run(addr string) error {
	app.scheduler()
	app.backgroundTask()
	app.middleware(
		middleware.GinMiddleware,
		middleware.Sentry,
	)
	app.router(
		router.Common,
		router.Auth,
		router.OAuth2,
		router.Users,
		router.Docs,
	)
	return app.engine.Run(addr)
}
