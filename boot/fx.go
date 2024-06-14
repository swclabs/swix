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

var _FxDataLayer = fx.Options(
	fx.Provide(
		config.LoadEnv,
		db.CreateConnection,
	),
	repository.FxModule,
)

var _FxBusinessLogic = fx.Options(
	service.FxModule,
)

var FxRestModule = fx.Options(
	_FxDataLayer,     // data layer constructor
	_FxBusinessLogic, // business logic constructor
	http.FxModule,    // presenter layer constructor
)

var FxWorkerModule = fx.Options(
	_FxDataLayer,     // data layer constructor
	_FxBusinessLogic, // business logic constructor
	workers.FxModule, // presenter layer constructor
)
