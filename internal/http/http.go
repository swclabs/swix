package http

import (
	"errors"
	"log"
	"os"
	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/internal/http/middleware"
	"swclabs/swipecore/internal/http/router"
	"swclabs/swipecore/pkg/sentry"

	"github.com/labstack/echo/v4"
)

type IServer interface {
	_BackgroundTask(tasks ...func())
	_InitMiddleware()
	_LoggerWriter(*os.File)
	middleware(mdws ...func(*echo.Echo))
	Connect(routers router.IRouter)
	Run(string) error
}

type _Server struct {
	engine *echo.Echo
}

var _ IServer = &_Server{}

func NewServer(common *router.Common, swaggerdocs *router.Docs) IServer {
	sentry.Init()
	server := &_Server{
		engine: echo.New(),
	}
	server._InitMiddleware()
	server.Connect(common)
	server.Connect(swaggerdocs)
	return server
}

func (server *_Server) middleware(mdws ...func(*echo.Echo)) {
	for _, m := range mdws {
		m(server.engine)
	}
}

func (server *_Server) _BackgroundTask(tasks ...func()) {
	for _, t := range tasks {
		go t()
	}
}

func (server *_Server) _LoggerWriter(file *os.File) {
	middleware.Logger(file, server.engine)
}

func (server *_Server) _InitMiddleware() {
	server.middleware(
		middleware.BaseSetting,
		middleware.CookieSetting,
		middleware.Sentry,
	)
}

func (server *_Server) Connect(routers router.IRouter) {
	routers.Routers(server.engine)
}

func (server *_Server) Run(addr string) error {
	if config.StageStatus != "dev" {
		const filePath = "api.log"
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return errors.New("error opening file: " + err.Error())
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(file)
		server._LoggerWriter(file)
	}
	return server.engine.Start(addr)
}
