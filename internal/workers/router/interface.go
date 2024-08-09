// Package router define tasks - queue
package router

import "swclabs/swix/internal/workers/server"

// IRouter interface for router objects
type IRouter interface {
	Register(eng server.IWorker)
}
