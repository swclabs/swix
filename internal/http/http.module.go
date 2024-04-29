package http

import (
	"swclabs/swipecore/internal/http/router"
)

func CommonModule(server IServer) {
	server.connect(router.New(router.TypeDocs))
	server.connect(router.New(router.TypeCommon))
}

func AccountManagementModule(server IServer) {
	server.connect(router.New(router.TypeAccountManagement))
}

func ProductsModule(server IServer) {
	server.connect(router.New(router.TypeProducts))
}

func PostsModule(server IServer) {
	server.connect(router.New(router.TypePosts))
}
