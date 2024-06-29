package repository

import (
	"swclabs/swipecore/internal/core/repository/accounts"
	"swclabs/swipecore/internal/core/repository/addresses"
	"swclabs/swipecore/internal/core/repository/carts"
	"swclabs/swipecore/internal/core/repository/categories"
	"swclabs/swipecore/internal/core/repository/collections"
	"swclabs/swipecore/internal/core/repository/inventory"
	"swclabs/swipecore/internal/core/repository/orders"
	"swclabs/swipecore/internal/core/repository/products"
	"swclabs/swipecore/internal/core/repository/suppliers"
	"swclabs/swipecore/internal/core/repository/users"

	"go.uber.org/fx"
)

var FxModule = fx.Options(
	fx.Provide(
		users.New,
		accounts.New,
		addresses.New,
		categories.New,
		products.New,
		suppliers.New,
		inventory.New,
		collections.New,
		orders.New,
		carts.New,
	),
)
