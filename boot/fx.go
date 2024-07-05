package boot

import (
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
	ProdMode
	DebugMode
)

var (
	_FxDataLayer      = fx.Options(fx.Provide(db.New), repository.FxModule)
	_FxBusinessLogic  = fx.Options(service.FxModule)
	_FxPresenterLayer = fx.Provide()
	_Logger           = fx.Provide()
)

func PrepareFor(flag int) {
	if flag&RestAPI != 0 {
		_FxPresenterLayer = http.FxModule
	}
	if flag&WorkerConsume != 0 {
		_FxPresenterLayer = workers.FxModule
	}
	if flag&DebugMode != 0 {
		_Logger = fx.Provide()
	}
	if flag&ProdMode != 0 {
		_Logger = fx.NopLogger
	}
}

func FxModule() fx.Option {
	return fx.Options(
		_FxDataLayer,      // data layer constructor
		_FxBusinessLogic,  // business logic constructor
		_FxPresenterLayer, // presenter layer constructor
		_Logger,
	)
}
