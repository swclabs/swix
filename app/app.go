package app

import (
	"context"
	"log"
	"swclabs/swipex/pkg/infra/blob"
	"swclabs/swipex/pkg/infra/cache"
	"swclabs/swipex/pkg/infra/db"
	"swclabs/swipex/pkg/lib/logger"

	"go.uber.org/fx"
)

type IApplication interface {
	Run() error
}

var module = fx.Provide()

func Repos[Fn any](reposConstructor Fn) Fn {
	module = fx.Options(module, fx.Provide(reposConstructor))
	return reposConstructor
}

func Service[Fn any](serviceConstructor Fn) Fn {
	module = fx.Options(module, fx.Provide(serviceConstructor))
	return serviceConstructor
}

func Router[Fn any](routersConstructor Fn) Fn {
	module = fx.Options(module, fx.Provide(routersConstructor))
	return routersConstructor
}

func Controller[Fn any](controllerConstructor Fn) Fn {
	module = fx.Options(module, fx.Provide(controllerConstructor))
	return controllerConstructor
}

func _main(lc fx.Lifecycle, app IApplication) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			logger.Info("Server starting")
			go func() {
				log.Fatal(app.Run())
			}()
			return nil
		},
		OnStop: func(_ context.Context) error {
			logger.Info("Server stopping")
			return nil
		},
	})
}

func Builder(constructor interface{}) *app {
	module = fx.Options(module, fx.Provide(constructor))
	return &app{
		core: fx.New(fx.Provide(blob.New, db.New, cache.New), module,
			fx.Invoke(_main),
		),
	}
}

type app struct {
	core *fx.App
}

func (a *app) Run() error {
	a.core.Run()
	return a.core.Err()
}
