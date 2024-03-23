package http

import (
	"errors"
	"log"
	"os"

	"github.com/swclabs/swipe-api/internal/http/middleware"

	"github.com/swclabs/swipe-api/internal/config"
	"github.com/swclabs/swipe-api/pkg/sentry"
	"github.com/swclabs/swipe-api/pkg/tools/worker"
)

var _ IServer = &_Server{}

func init() {
	sentry.Init()
	// mailers.Config(config.Email, config.EmailAppPassword)
	worker.SetBroker(config.RedisHost, config.RedisPort, config.RedisPassword)
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

func (server *_Server) Bootstrap(fn ...func(server IServer)) {
	for _, _fn := range fn {
		_fn(server)
	}
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
