// Package boot define all specific function to start
package boot

import (
	"swclabs/swipecore/internal/core/repository"
	"swclabs/swipecore/internal/core/service"
	"swclabs/swipecore/internal/wapi"
	"swclabs/swipecore/internal/workers"
	"swclabs/swipecore/pkg/infra/blob"
	"swclabs/swipecore/pkg/infra/db"

	"go.uber.org/fx"
)

const (
	// RestAPI flag build web api
	RestAPI = 1 << iota
	// WorkerConsume flag build worker
	WorkerConsume
	// ProdMode build with production mode
	ProdMode
	// DebugMode build with developer mode
	DebugMode
)

var (
	_FxInfrastructure = fx.Provide(blob.New, db.New)
	_FxDataLayer      = fx.Options(repository.FxModule)
	_FxBusinessLogic  = fx.Options(service.FxModule)
	_FxPresenterLayer = fx.Provide()
	_Logger           = fx.Provide()
)

// PrepareFor enable build web api or worker consume
func PrepareFor(flag int) {
	if flag&RestAPI != 0 {
		_FxPresenterLayer = wapi.FxModule
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

func fxModule() fx.Option {
	return fx.Options(
		_FxInfrastructure,
		_FxDataLayer,      // data layer constructor
		_FxBusinessLogic,  // business logic constructor
		_FxPresenterLayer, // presenter layer constructor
		_Logger,
	)
}
