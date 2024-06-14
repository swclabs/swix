package controller

import (
	"go.uber.org/fx"
)

var FxModule = fx.Options(
	fx.Provide(
		NewPosts,
		NewProducts,
		NewAccountManagement,
		NewCommon,
	),
)
