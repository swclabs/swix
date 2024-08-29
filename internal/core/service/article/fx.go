package article

import (
	"swclabs/swix/internal/core/repository/collections"
	"swclabs/swix/pkg/infra/blob"
	"swclabs/swix/pkg/infra/db"

	"go.uber.org/fx"
)

var FxModule = fx.Option(
	fx.Provide(
		blob.New,
		db.New,
		collections.New,
		New,
	),
)
