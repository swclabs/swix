package api

import (
	"time"

	"swclabs/swiftcart/api/middleware"
	"swclabs/swiftcart/api/router"
	"swclabs/swiftcart/internal/config"
	"swclabs/swiftcart/internal/service"
	"swclabs/swiftcart/internal/tasks"
	"swclabs/swiftcart/pkg/sentry"
	"swclabs/swiftcart/pkg/x/job"
	"swclabs/swiftcart/pkg/x/mailers"
	"swclabs/swiftcart/pkg/x/worker"
)

func init() {
	sentry.Init()
	mailers.Config(config.Email, config.EmailAppPassword)
	worker.SetBroker(config.RedisHost, config.RedisPort, config.RedisPassword)
}

func (s *Server) Scheduler() {
	newJob := job.New()
	go newJob.Scheduler(service.Ping, 5*time.Second)

	newJob.Launch()
}

func (s *Server) Run(addr string) error {
	s.Scheduler()
	s.backgroundTask()
	s.middleware(
		middleware.GinMiddleware,
		middleware.Sentry,
	)
	s.router(
		router.Common,
		router.Auth,
		router.OAuth2,
		router.Users,
		router.Docs,
	)
	return s.engine.Run(addr)
}

func (w *Worker) Run() error {
	w.engine.HandleFunctions(tasks.Path())
	return w.engine.Run()
}
