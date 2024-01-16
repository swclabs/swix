package delivery

import (
	"swclabs/swiftcart/internal/http"
)

type HttpServer struct {
	Server *http.Server
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		Server: http.New(),
	}
}

func (adapter *HttpServer) ListenOn(port string) error {
	return adapter.Server.Run(port)
}
