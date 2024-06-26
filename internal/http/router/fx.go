package router

import (
	"go.uber.org/fx"
)

var FxModule = fx.Options(
	fx.Provide(
		New,
		NewProducts,
		NewAccountManagement,
		NewPosts,
		NewPurchase,
	),
)
