
package apis

import (
	"swclabs/swix/boot"
	"swclabs/swix/internal/apis/router"
	"swclabs/swix/internal/apis/server"
)

func NewAPIServer(
	base router.IBaseRouter,
	products router.IProducts,
	manager router.IManager,
	article router.IArticle,
	purchase router.IPurchase,
	classify router.IClassify,
	paydeli router.IPaydeliver,
) boot.ICore {
	mux := server.NewServeMux()
	mux.Handle(base)
	mux.Handle(products)
	mux.Handle(manager)
	mux.Handle(article)
	mux.Handle(purchase)
	mux.Handle(classify)
	mux.Handle(paydeli)
	server := server.New(mux)
	return server
}
