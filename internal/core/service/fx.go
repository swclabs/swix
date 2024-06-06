package service

import (
	"swclabs/swipecore/internal/core/service/accountmanagement"
	"swclabs/swipecore/internal/core/service/common"
	"swclabs/swipecore/internal/core/service/posts"
	"swclabs/swipecore/internal/core/service/products"
	"swclabs/swipecore/pkg/lib/worker"

	"go.uber.org/fx"
)

var FxModule = fx.Options(
	fx.Provide(
		worker.NewClient,

		common.New,
		products.New,
		accountmanagement.New,
		posts.New,
	),
)
