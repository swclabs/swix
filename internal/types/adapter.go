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
// │  ...   │  │ webapi.IServer │  │ worker.Writer  │       │   uber/fx app  │
// │        │  │                │  │                │       │                │
// └────────┘  └────────────────┘  └────────────────┘       └────────────────┘
// The outer layer is the Adapter (types.Adapter), implemented in the
// main.adapter.go.We use the IServer in the boot directory to connect
// with the IServer of the webapi and worker directory through the
// Adapter interface.

// Package types define internal types
package types

// IAdapter interface for adapter objects
type IAdapter interface {
	Run(addr string) error
	StartWorker(concurrency int) error
}
