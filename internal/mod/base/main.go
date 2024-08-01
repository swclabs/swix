package base

import (
	"net"
	"swclabs/swipecore/internal/mod/proto/base"
	"swclabs/swipecore/internal/types"
	"swclabs/swipecore/pkg/lib/logger"

	"google.golang.org/grpc"
)

var _ types.IAdapter = (*_BaseServer)(nil)

// New returns a new base server adapter.
func New(baseServer base.BaseServer) types.IAdapter {
	srv := grpc.NewServer(grpc.UnaryInterceptor(logger.GRPCLogger))

	base.RegisterBaseServer(srv, baseServer)

	return &_BaseServer{
		server: srv,
	}
}

// _BaseServer is the base server grpc implementation.
type _BaseServer struct {
	server *grpc.Server
}

// Run implements types.IAdapter.
func (srv *_BaseServer) Run(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	logger.GRPCServerINFO("BASE", addr)
	return srv.server.Serve(listener)
}
