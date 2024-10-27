package apis

import (
	"swclabs/swipex/app"
	"swclabs/swipex/internal/apis/container/article"
	"swclabs/swipex/internal/apis/container/authentication"
	"swclabs/swipex/internal/apis/container/classify"
	"swclabs/swipex/internal/apis/container/healthcheck"
	"swclabs/swipex/internal/apis/container/products"
	"swclabs/swipex/internal/apis/container/purchase"
	"swclabs/swipex/internal/apis/server"
)

func NewApp(
	base healthcheck.IRouter,
	article article.IRouter,
	purchase purchase.IRouter,
	classify classify.IRouter,
	auth authentication.IRouter,
	products products.IRouter,
) app.IApplication {
	mux := server.NewServeMux()
	mux.Handle(base)
	mux.Handle(article)
	mux.Handle(purchase)
	mux.Handle(classify)
	mux.Handle(auth)
	mux.Handle(products)
	return server.New(mux)
}
