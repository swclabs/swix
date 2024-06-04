package repository

import (
	"swclabs/swipecore/internal/core/repository/accounts"
	"swclabs/swipecore/internal/core/repository/addresses"
	"swclabs/swipecore/internal/core/repository/categories"
	"swclabs/swipecore/internal/core/repository/collections"
	"swclabs/swipecore/internal/core/repository/products"
	"swclabs/swipecore/internal/core/repository/suppliers"
	"swclabs/swipecore/internal/core/repository/users"
	"swclabs/swipecore/internal/core/repository/warehouse"

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
		warehouse.New,
		collections.New,
	),
)
