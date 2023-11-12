package api

import (
	"example/komposervice/api/middleware"
	"example/komposervice/api/router"
	"example/komposervice/internal/config"
	"example/komposervice/internal/service"
	"example/komposervice/internal/tasks"
	"example/komposervice/pkg/lib/job"
	"example/komposervice/pkg/lib/mailers"
	"example/komposervice/pkg/lib/worker"
	"example/komposervice/pkg/sentry"
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
		router.Auth,
		router.Docs,
	)
	return s.engine.Run(addr)
}

func (w *Worker) Run() error {
	w.engine.HandleFunctions(tasks.Path())
	return w.engine.Run()
}
