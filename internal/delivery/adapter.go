package delivery

import (
	"swclabs/swiftcart/internal/http"
)

func Adapter(server *http.Server) {
	server.InitAccountManagement()
}

func AccountManagementAdapter(server *http.Server) {
	server.InitAccountManagement()
}
