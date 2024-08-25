package boot

import (
	"swclabs/swix/internal/types"
)

func NewGate() IServer {
	return &_Gate{address: "localhost:8080"}
}

type _Gate struct {
	address string
}

// Connect implements IServer.
func (gate *_Gate) Connect(adapter types.IAdapter) error {
	return adapter.Run(gate.address)
}
