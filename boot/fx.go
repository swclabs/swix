// Package boot define all specific function to start
package boot

import (
	"swclabs/swix/internal/core/repository"
	"swclabs/swix/internal/core/service"
	"swclabs/swix/internal/webapi"
	"swclabs/swix/internal/workers"
	"swclabs/swix/pkg/infra/blob"
	"swclabs/swix/pkg/infra/cache"
	"swclabs/swix/pkg/infra/db"

	"go.uber.org/fx"
)

const (
	// WebAPI flag build web api
	WebAPI = 1 << iota
	// Worker flag build worker
	Worker
	// ProdMode build with production mode
	ProdMode
	// DebugMode build with developer mode
	DebugMode
)

var (
	_FxInfrastructure = fx.Provide(blob.New, db.New, cache.New)
	_FxDataLayer      = fx.Options(repository.FxModule)
	_FxBusinessLogic  = fx.Options(service.FxModule)
	_FxPresenterLayer = fx.Provide()
	_Logger           = fx.Provide()
)

// PrepareFor enable build web api or worker consume
func PrepareFor(flag int) {
	if flag&WebAPI != 0 {
		_FxPresenterLayer = webapi.FxModule
	}
	if flag&Worker != 0 {
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
