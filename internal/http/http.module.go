package http

import (
	"swclabs/swipe-api/internal/http/router"
)

func CommonModule(server IServer) {
	server.router(
		router.Common,
		router.Docs,
	)
}

func AccountManagementModule(server IServer) {
	var accountManagement = router.NewAccountManagement()
	server.router(
		accountManagement.Users,
		accountManagement.Auth,
		accountManagement.OAuth2,
	)
}

func ProductManagementModule(server IServer) {
	var productManagement = router.NewProductManagement()
	server.router(
		productManagement.Category,
		productManagement.Product,
		productManagement.Newsletter,
	)
}

func ProductsModule(server IServer) {
	var products = router.NewProducts()
	server.router(
		products.Common,
	)
}
