package api

import (
	"example/komposervice/internal/config"
	"example/komposervice/internal/tasks"
	"example/komposervice/pkg/lib/worker"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	if config.StageStatus != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}
	return &Server{
		engine: gin.Default(),
	}
}

func (s *Server) middleware(mdws ...func(*gin.Engine)) {
	for _, m := range mdws {
		m(s.engine)
	}
}

func (s *Server) backgroundTask(tasks ...func()) {
	for _, t := range tasks {
		go t()
	}
}

func (s *Server) router(routers ...func(*gin.Engine)) {
	for _, r := range routers {
		r(s.engine)
	}
}

type Worker struct {
	engine *worker.Engine
}

func NewWorker(concurrency int) *Worker {
	return &Worker{
		engine: worker.NewServer(concurrency, tasks.Queue()),
	}
}
