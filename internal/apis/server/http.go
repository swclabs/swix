package server

import (
	"errors"
	"fmt"
	"log"
	"os"
	"swclabs/swix/internal/apis/middleware"
	"swclabs/swix/internal/config"
	"swclabs/swix/pkg/infra/sentry"

	"github.com/labstack/echo/v4"
)

// IServer interface represents all server method
type IServer interface {
	Run() error
}

var _ IServer = &_Server{}

// New creates a new instance of the Server
func New(mux IMux) IServer {
	sentry.Init()
	server := &_Server{
		mux:    mux,
		engine: echo.New(),
	}
	server.initMiddleware()
	return server
}

type _Server struct {
	mux    IMux
	engine *echo.Echo
}

func (server *_Server) loggerWriter(file *os.File) {
	middleware.Logger(file, server.engine)
}

func (server *_Server) initMiddleware() {
	middleware.BaseSetting(server.engine)
	middleware.CookieSetting(server.engine)
	middleware.Sentry(server.engine)
}

func (server *_Server) Run() error {
	server.mux.ServeHTTP(server.engine)
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
	return server.engine.Start(fmt.Sprintf("%s:%s", config.Host, config.Port))
}
