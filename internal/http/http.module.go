package http

import (
	"swclabs/swipe-api/internal/helper/resolver"
	"swclabs/swipe-api/internal/http/router"
)

func CommonModule(server IServer) {
	server.router(
		router.Common,
		router.Docs,
	)
}

func AccountManagementModule(server IServer) {
	server._BackgroundTask(func() {
		resolver.StarUserImageHandler(5)
	})
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
	)
}
