// Package types defines the interface of gateway and grpc server.
package types

// IGrpcServer is the interface for gRPC server.
type IGrpcServer interface {
	Port() int
	Serve() error
}
