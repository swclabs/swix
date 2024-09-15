// Package repos implement the repos layer
package repos

import (
	"swclabs/swix/internal/core/repos/accounts"
	"swclabs/swix/internal/core/repos/addresses"
	"swclabs/swix/internal/core/repos/carts"
	"swclabs/swix/internal/core/repos/categories"
	"swclabs/swix/internal/core/repos/collections"
	"swclabs/swix/internal/core/repos/comments"
	"swclabs/swix/internal/core/repos/deliveries"
	"swclabs/swix/internal/core/repos/inventories"
	"swclabs/swix/internal/core/repos/orders"
	"swclabs/swix/internal/core/repos/products"
	"swclabs/swix/internal/core/repos/specifications"
	"swclabs/swix/internal/core/repos/suppliers"
	"swclabs/swix/internal/core/repos/users"

	"go.uber.org/fx"
)

// FxModule module of package repos
var FxModule = fx.Options(
	fx.Provide(
		users.Init,
		accounts.Init,
		addresses.Init,
		categories.Init,
		products.Init,
		suppliers.Init,
		inventories.Init,
		collections.Init,
		orders.Init,
		carts.Init,
		specifications.Init,
		deliveries.Init,
		comments.Init,
	),
)
