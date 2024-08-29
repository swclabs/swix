package nodes

// IGrpcServer is the interface for gRPC server.
type IGrpcServer interface {
	Port() int
	Serve() error
}
