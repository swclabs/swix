package api

import (
	"swclabs/swiftcart/internal/config"
	"swclabs/swiftcart/internal/tasks"
	"swclabs/swiftcart/pkg/x/worker"

	"github.com/gin-gonic/gin"
)

type IServer interface {
	middleware(...func(*gin.Engine))
	backgroundTask(...func())
	router(...func(*gin.Engine))
	Scheduler()
	Run(string) error
}

type IWorker interface {
	Run() error
}

type Server struct {
	engine *gin.Engine
}

func NewServer() IServer {
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

func NewWorker(concurrency int) IWorker {
	return &Worker{
		engine: worker.NewServer(concurrency, tasks.Queue()),
	}
}
