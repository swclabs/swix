// Package port define port
package port

import (
	"fmt"
)

const (
	// Greeter is the port for Greeter service.
	Greeter = 1000 << iota
)

const (
	// Gateway is the port for the gateway.
	Gateway = 8080
)

// Addr is the address of the services.
var Addr = map[int]string{
	Greeter: fmt.Sprintf("%s:%d", "localhost", Greeter),
}
