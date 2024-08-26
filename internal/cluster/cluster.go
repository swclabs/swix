// Package cluster defines the gateway of the service.
// grpc server cluster is the server that serves the grpc service.
package cluster

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"swclabs/swix/internal/cluster/nodes"
	"swclabs/swix/internal/cluster/port"
	"swclabs/swix/internal/cluster/proto/greeter"
	"swclabs/swix/pkg/lib/logger"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type (
	// registerHandler defines the data type as a registerHandler function,
	// same type as the function generated when registering the gateway
	registerHandler func(
		ctx context.Context,
		mux *runtime.ServeMux,
		endpoint string,
		opts []grpc.DialOption) (err error)

	// handler is a map that contains ports and corresponding registerHandler functions
	// port is of type int, equivalent address is [::port].
	// registerHandler is the data type of the corresponding registerHandler
	// function for the registered service.
	//
	// Example:
	//
	// func greeter.RegisterGreeterHandlerFromEndpoint(
	//		ctx context.Context,
	//   	mux *runtime.ServeMux,
	//      endpoint string, opts []grpc.DialOption) (err error)
	//
	// is used to register with the gateway through runtime.ServeMux corresponding to port.Greeter.
	handler map[int]registerHandler
)

// ICluster is the interface for the cluster.
type ICluster interface {
	ServeNode(flag int)
	ServeMux(ctx context.Context, mux *runtime.ServeMux, opts []grpc.DialOption) error
}

// cluster is the data type representing a cluster that contains registered services
type cluster struct {
	handler handler
	node    []nodes.IGrpcServer
}

// ServeMux registers services with the gateway through runtime.ServeMux
// each service is registered with a corresponding port
func (c *cluster) ServeMux(ctx context.Context, mux *runtime.ServeMux, opts []grpc.DialOption) error {
	// iterate through the handler map, register each service with the gateway
	for _port, registerHandler := range c.handler {
		addr := port.Addr[_port]
		// register through runtime.ServeMux with the port and grpc options
		if err := registerHandler(ctx, mux, addr, opts); err != nil {
			return err
		}
	}
	return nil
}

// ServeNode starts the gRPC cluster server.
// flag is the bitwise OR of the ports to be started.
func (c *cluster) ServeNode(flag int) {
	logger.Info("Start Cluster with flowing setting:")
	for idx, s := range c.node {
		if flag&s.Port() != 0 {
			go func() {
				logger.Info(fmt.Sprintf("node #%d ===> listen on [::]%s",
					idx,
					logger.Green.Add(strconv.Itoa(s.Port()))),
				)
				if err := s.Serve(); err != nil {
					log.Fatalf("Failed to serve: %v", err)
				}
			}()
		}
	}
}

// New creates a new cluster with registered services, initialized through the fx framework
func New(
	greeterNode nodes.IGreeter,
) ICluster {
	return &cluster{
		handler: handler{
			port.Greeter: registerHandler(greeter.RegisterGreeterHandlerFromEndpoint),
		},
		node: []nodes.IGrpcServer{
			greeterNode,
		},
	}
}
