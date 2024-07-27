// Package service implements usecase layer
package service

import (
	"swclabs/swipecore/internal/core/service/classify"
	"swclabs/swipecore/internal/core/service/common"
	"swclabs/swipecore/internal/core/service/manager"
	"swclabs/swipecore/internal/core/service/posts"
	"swclabs/swipecore/internal/core/service/products"
	"swclabs/swipecore/internal/core/service/purchase"

	"go.uber.org/fx"
)

// FxModule module of package service
var FxModule = fx.Options(
	fx.Provide(
		common.New,
		products.New,
		manager.New,
		posts.New,
		purchase.New,
		classify.New,
	),
)
