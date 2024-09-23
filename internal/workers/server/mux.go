package server

import (
	"swclabs/swix/pkg/lib/worker"
)

type IMux interface {
	Handle(router IRouter)
	Serve(enginer worker.IEngine)
}

func NewServeMux() IMux {
	return &Mux{}
}

type Mux struct {
	router []IRouter
}

// Handle implements IMux.
func (m *Mux) Handle(router IRouter) {
	m.router = append(m.router, router)
}

// Serve implements IMux.
func (m *Mux) Serve(enginer worker.IEngine) {
	for _, r := range m.router {
		r.Register(enginer)
	}
}
