// Package repository implement the repository layer
package repository

import (
	"swclabs/swix/internal/core/repository/accounts"
	"swclabs/swix/internal/core/repository/addresses"
	"swclabs/swix/internal/core/repository/carts"
	"swclabs/swix/internal/core/repository/categories"
	"swclabs/swix/internal/core/repository/collections"
	"swclabs/swix/internal/core/repository/inventories"
	"swclabs/swix/internal/core/repository/orders"
	"swclabs/swix/internal/core/repository/products"
	"swclabs/swix/internal/core/repository/specifications"
	"swclabs/swix/internal/core/repository/suppliers"
	"swclabs/swix/internal/core/repository/users"

	"go.uber.org/fx"
)

// FxModule module of package repository
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
	),
)
