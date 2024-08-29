// Package nodes defines the grpc server.
package nodes

import (
	"fmt"
	"net"
	"swclabs/swix/internal/cluster/port"
	"swclabs/swix/internal/cluster/proto/greeter"
	"swclabs/swix/pkg/lib/logger"

	"google.golang.org/grpc"
)

type _Greeter struct {
	grpcServer *grpc.Server
	service    greeter.GreeterServer
}

// IGreeter is the interface for Greeter service.
type IGreeter interface {
	IGrpcServer
}

// NewGreeter creates a new Greeter service.
func NewGreeter(srv greeter.GreeterServer) IGreeter {
	return &_Greeter{
		grpcServer: grpc.NewServer(
			grpc.UnaryInterceptor(logger.Logger),
		),
		service: srv,
	}
}

// Port implements IGreeter.
func (g *_Greeter) Port() int {
	return port.Greeter
}

// Serve implements IGreeter.
func (g *_Greeter) Serve() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port.Greeter))
	if err != nil {
		return err
	}

	// Start gRPC server here
	greeter.RegisterGreeterServer(g.grpcServer, g.service)
	return g.grpcServer.Serve(listener)
}
