package webapi

import (
	"errors"
	"fmt"
	"log"
	"os"
	"swclabs/swix/internal/config"
	"swclabs/swix/internal/webapi/middleware"
	"swclabs/swix/internal/webapi/router"
	"swclabs/swix/pkg/infra/sentry"

	"github.com/labstack/echo/v4"
)

type _Server struct {
	engine *echo.Echo
}

var _ IServer = &_Server{}

// NewServer creates a new instance of the Server
func NewServer(base router.IBase) IServer {
	sentry.Init()
	server := &_Server{
		engine: echo.New(),
	}
	server.initMiddleware()
	server.Connect(base)
	return server
}

func (server *_Server) middleware(mdws ...func(*echo.Echo)) {
	for _, m := range mdws {
		m(server.engine)
	}
}

func (server *_Server) backgroundTask(tasks ...func()) {
	for _, t := range tasks {
		go t()
	}
}

func (server *_Server) loggerWriter(file *os.File) {
	middleware.Logger(file, server.engine)
}

func (server *_Server) initMiddleware() {
	server.middleware(
		middleware.BaseSetting,
		middleware.CookieSetting,
		middleware.Sentry,
	)
}

func (server *_Server) Routes() []string {
	var path = make([]string, 0)
	for _, route := range server.engine.Routes() {
		if route != nil {
			path = append(path,
				fmt.Sprintf("[%s]    %s \n", route.Method, route.Path))
		}
	}
	return path
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
		server.loggerWriter(file)
	}
	return server.engine.Start(addr)
}
