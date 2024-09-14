// Package router implements the router interface
package router

import (
	"go.uber.org/fx"
)

// FxModule module of package router
var FxModule = fx.Options(
	fx.Provide(
		New,
		NewProducts,
		NewArticle,
		NewPurchase,
		NewManager,
		NewClassify,
		NewPaydeliver,
	),
)
