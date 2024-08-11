// Package service implements usecase layer
package service

import (
	"swclabs/swix/internal/core/service/article"
	"swclabs/swix/internal/core/service/base"
	"swclabs/swix/internal/core/service/classify"
	"swclabs/swix/internal/core/service/manager"
	"swclabs/swix/internal/core/service/paydeliver"
	"swclabs/swix/internal/core/service/products"
	"swclabs/swix/internal/core/service/purchase"

	"go.uber.org/fx"
)

// FxModule module of package service
var FxModule = fx.Options(
	fx.Provide(
		base.New,
		products.New,
		manager.New,
		article.New,
		purchase.New,
		classify.New,
		paydeliver.New,
	),
)
