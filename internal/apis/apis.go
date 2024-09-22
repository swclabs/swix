package apis

import (
	"swclabs/swix/app"
	"swclabs/swix/internal/apis/container/article"
	"swclabs/swix/internal/apis/container/base"
	"swclabs/swix/internal/apis/container/classify"
	"swclabs/swix/internal/apis/container/manager"
	"swclabs/swix/internal/apis/container/products"
	"swclabs/swix/internal/apis/container/purchase"
	"swclabs/swix/internal/apis/server"
)

func NewApp(
	base base.IRouter,
	article article.IRouter,
	purchase purchase.IRouter,
	classify classify.IRouter,
	manager manager.IRouter,
	products products.IRouter,
) app.IApplication {
	mux := server.NewServeMux()
	mux.Handle(base)
	mux.Handle(article)
	mux.Handle(purchase)
	mux.Handle(classify)
	mux.Handle(manager)
	mux.Handle(products)
	return server.New(mux)
}
