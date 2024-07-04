package boot

import (
	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/internal/core/repository"
	"swclabs/swipecore/internal/core/service"
	"swclabs/swipecore/internal/http"
	"swclabs/swipecore/internal/workers"
	"swclabs/swipecore/pkg/db"

	"go.uber.org/fx"
)

const (
	RestAPI       = 1 << iota // 0001
	WorkerConsume             // 0010
)

var (
	_FxDataLayer      = fx.Options(fx.Provide(config.LoadEnv, db.New), repository.FxModule)
	_FxBusinessLogic  = fx.Options(service.FxModule)
	_FxPresenterLayer = http.FxModule
)

func PrepareFor(flag int) {
	if flag&RestAPI != 0 {
		_FxPresenterLayer = http.FxModule
	}
	if flag&WorkerConsume != 0 {
		_FxPresenterLayer = workers.FxModule
	}
}

func FxModule() fx.Option {
	return fx.Options(
		_FxDataLayer,      // data layer constructor
		_FxBusinessLogic,  // business logic constructor
		_FxPresenterLayer, // presenter layer constructor
	)
}
