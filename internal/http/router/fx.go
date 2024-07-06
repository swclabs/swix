package router

import (
	"go.uber.org/fx"
)

var FxModule = fx.Options(
	fx.Provide(
		NewDocs,
		NewCommon,
		NewProducts,
		NewAccountManagement,
		NewPosts,
		NewPurchase,
	),
)
