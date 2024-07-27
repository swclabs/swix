// Package router implements the router interface
package router

import (
	"go.uber.org/fx"
)

// FxModule module of package router
var FxModule = fx.Options(
	fx.Provide(
		NewDocs,
		NewCommon,
		NewProducts,
		NewPosts,
		NewPurchase,
		NewManager,
		NewClassify,
	),
)
