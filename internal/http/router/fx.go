package router

import (
	"go.uber.org/fx"
)

var FxModule = fx.Options(
	fx.Provide(
		NewProducts,
		NewCommon,
		NewDocs,
		NewAccountManagement,
		NewPosts,
	),
)
