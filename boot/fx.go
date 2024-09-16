// Package boot define all specific function to start
package boot

import (
	apis "swclabs/swix/internal/apis/server"
	"swclabs/swix/internal/core/repos"
	"swclabs/swix/internal/core/service"
	workers "swclabs/swix/internal/workers/server"
	"swclabs/swix/pkg/infra/blob"
	"swclabs/swix/pkg/infra/cache"
	"swclabs/swix/pkg/infra/db"

	"go.uber.org/fx"
)

const (
	// APIs flag build web api
	APIs = 1 << iota
	// Worker flag build worker
	Worker
	// ProdMode build with production mode
	ProdMode
	// DebugMode build with developer mode
	DebugMode
)

var (
	fxInfrastructure = fx.Provide(blob.New, db.New, cache.New)
	fxDataLayer      = fx.Options(repos.FxModule)
	fxBusinessLogic  = fx.Options(service.FxModule)
	fxPresenterLayer = fx.Provide()
	fxLogger         = fx.Provide()
)

func fxModule(flag int) fx.Option {
	switch {
	case flag&APIs != 0:
		fxPresenterLayer = apis.FxModule
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
