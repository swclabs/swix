package app

import (
	"time"

	"swclabs/swiftcart/app/middleware"
	"swclabs/swiftcart/app/router"
	"swclabs/swiftcart/internal/config"
	"swclabs/swiftcart/internal/service"
	_ "swclabs/swiftcart/internal/tasks/plugin"
	"swclabs/swiftcart/pkg/job"
	"swclabs/swiftcart/pkg/mailers"
	"swclabs/swiftcart/pkg/sentry"
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