package app

import (
	"swclabs/swiftcart/internal/config"

	"github.com/gin-gonic/gin"
)

type IApp interface {
	middleware(...func(*gin.Engine))
	backgroundTask(...func())
	router(...func(*gin.Engine))
	scheduler()
	Run(string) error
}

type App struct {
	engine *gin.Engine
}

func New() *App {
	if config.StageStatus != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}
	return &App{
		engine: gin.Default(),
	}
}

func (app *App) middleware(mdws ...func(*gin.Engine)) {
	for _, m := range mdws {
		m(app.engine)
	}
}

func (app *App) backgroundTask(tasks ...func()) {
	for _, t := range tasks {
		go t()
	}
}

func (app *App) router(routers ...func(*gin.Engine)) {
	for _, r := range routers {
		r(app.engine)
	}
}
