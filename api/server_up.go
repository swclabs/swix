package api

import (
	"example/swiftcart/api/middleware"
	"example/swiftcart/api/router"
	"example/swiftcart/internal/config"
	"example/swiftcart/internal/service"
	"example/swiftcart/internal/tasks"
	"example/swiftcart/pkg/lib/job"
	"example/swiftcart/pkg/lib/mailers"
	"example/swiftcart/pkg/lib/worker"
	"example/swiftcart/pkg/sentry"
	"time"
)

func init() {
	sentry.Init()
	mailers.Config(config.Email, config.EmailAppPassword)
	worker.SetBroker(config.RedisHost, config.RedisPort, config.RedisPassword)
}

func (s *Server) Scheduler() {
	j := job.New()
	j.Scheduler(service.Ping, 5*time.Second)

	if err := j.Launch(); err != nil {
		panic(err)
	}
}

func (s *Server) Run(addr string) error {
	s.backgroundTask(
		s.Scheduler,
	)
	s.middleware(
		middleware.GinMiddleware,
		middleware.Sentry,
	)
	s.router(
		router.Common,
		// router.Auth,
		router.Docs,
	)
	return s.engine.Run(addr)
}

func (w *Worker) Run() error {
	w.engine.HandleFunctions(tasks.Path())
	return w.engine.Run()
}
