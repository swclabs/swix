//          ┌────────────────┐                              ┌────────────────┐
//          │                │     boot.IServer.Connect()   │                │
//     ┌───►│ types.Adapter  ├─────────────────────────────►│  boot.IServer  │
//     │    │                │                              │                │
//     │    └────────────────┘◄─────────────┐               └───────┬────────┘
//     │               ▲                    │                       │
//     │inheritance    │                    │                       │
//     │               │inheritance         │inheritance            │invoke
//     │               │                    │                       │
//     │               │                    │                       ▼
// ┌───┴────┐  ┌───────┴────────┐  ┌────────┴───────┐       ┌────────────────┐
// │        │  │                │  │                │       │                │
// │  ...   │  │   apis.IServer │  │ worker.Writer  │       │   uber/fx app  │
// │        │  │                │  │                │       │                │
// └────────┘  └────────────────┘  └────────────────┘       └────────────────┘
// The outer layer is the Adapter (types.Adapter), implemented in the
// main.adapter.go We use the IServer in the boot directory to connect
// with the IServer of the apis and worker directory through the
// Adapter interface.

// Package types define internal types
package types

// IAdapter interface for adapter objects
type IAdapter interface {
	// Run starts the server via the adapter
	// prop is the address of the server if you are using the apis or gRPC adapter
	// or the number of workers if you are using the worker adapter
	//
	// apis or gRPC Example:
	//
	//  adapter.Run("localhost:8080")
	//
	// Worker Example:
	//
	//  adapter.Run("10")
	Run(prop string) error
}
