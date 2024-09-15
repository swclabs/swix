// Package controller implements the controller interface
package controller

import (
	"go.uber.org/fx"
)

// FxModule module of package controller
var FxModule = fx.Options(
	fx.Provide(
		New,
		NewArticle,
		NewProducts,
		NewManager,
		NewPurchase,
		NewClassify,
		NewPaydeliver,
	),
)
