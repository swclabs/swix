package http

import (
	"github.com/swclabs/swipe-api/internal/helper/resolver"
	"github.com/swclabs/swipe-api/internal/http/router"
)

func CommonModule(server IServer) {
	server.router(
		router.Common,
		router.Docs,
	)
}

func AccountManagementModule(server IServer) {
	server.backgroundTask(func() {
		resolver.StartImageHandler(5)
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
	)
}
