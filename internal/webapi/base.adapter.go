package webapi

import (
	"swclabs/swix/internal/types"
	"swclabs/swix/internal/webapi/server"
)

// NewBaseAdapter returns a new adapter wrapping around the given server
func NewBaseAdapter(server server.IServer) types.IAdapter {
	adapter := &_Adapter{
		server: server,
	}
	return adapter
}
