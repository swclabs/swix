// Package service implements usecase layer
package service

import (
	"swclabs/swipecore/internal/core/service/accountmanagement"
	"swclabs/swipecore/internal/core/service/common"
	"swclabs/swipecore/internal/core/service/posts"
	"swclabs/swipecore/internal/core/service/products"
	"swclabs/swipecore/internal/core/service/purchase"

	"go.uber.org/fx"
)

var FxModule = fx.Options(
	fx.Provide(
		common.New,
		products.New,
		accountmanagement.New,
		posts.New,
		purchase.New,
	),
)
