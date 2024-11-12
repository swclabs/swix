package apis

import (
	"github.com/swclabs/swipex/app"
	"github.com/swclabs/swipex/internal/apis/container/article"
	"github.com/swclabs/swipex/internal/apis/container/authentication"
	"github.com/swclabs/swipex/internal/apis/container/classify"
	"github.com/swclabs/swipex/internal/apis/container/healthcheck"
	"github.com/swclabs/swipex/internal/apis/container/products"
	"github.com/swclabs/swipex/internal/apis/container/purchase"
	"github.com/swclabs/swipex/internal/apis/server"
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
