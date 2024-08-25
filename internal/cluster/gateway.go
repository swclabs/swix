// Package cluster define the gateway of the service.
// grpc server cluster is the server that serves the grpc service.
package cluster

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"swclabs/swix/internal/cluster/port"
	"swclabs/swix/pkg/lib/logger"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// RegisHandler is the register handler function for the gateway.
type RegisHandler func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

// IGateway is the interface for the gateway.
type IGateway interface {
	Connect(ctx context.Context, cluster ICluster) error
	ListenAndServe(addr string) error
}

// Gateway is the gateway of the service.
type Gateway struct {
	mux  *runtime.ServeMux
	opts []grpc.DialOption
}

// NewGateway creates a new gateway.
func NewGateway() IGateway {
	return &Gateway{
		mux:  runtime.NewServeMux(),
		opts: []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	}
}

// Connect connects the gateway to the cluster.
func (g *Gateway) Connect(ctx context.Context, cluster ICluster) error {
	return cluster.ServeMux(ctx, g.mux, g.opts)
}

// ListenAndServe start the gateway server.
func (g *Gateway) ListenAndServe(addr string) error {
	logger.Info(fmt.Sprintf("gRPC Gateway listen on [::]%s", logger.Cyan.Add(strconv.Itoa(port.Gateway))))
	return http.ListenAndServe(addr, g.mux)
}
