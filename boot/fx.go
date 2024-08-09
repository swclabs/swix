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
	fxInfrastructure = fx.Provide(blob.New, db.New, cache.New)
	fxDataLayer      = fx.Options(repository.FxModule)
	fxBusinessLogic  = fx.Options(service.FxModule)
	fxPresenterLayer = fx.Provide()
	fxLogger         = fx.Provide()
)

func fxModule(flag int) fx.Option {
	switch {
	case flag&WebAPI != 0:
		fxPresenterLayer = webapi.FxModule
	case flag&Worker != 0:
		fxPresenterLayer = workers.FxModule
	}
	switch {
	case flag&DebugMode != 0:
		fxLogger = fx.Provide()
	case flag&ProdMode != 0:
		fxLogger = fx.NopLogger
	}
	return fx.Options(
		fxInfrastructure,
		fxDataLayer,      // data layer constructor
		fxBusinessLogic,  // business logic constructor
		fxPresenterLayer, // presenter layer constructor
		fxLogger,
	)
}
