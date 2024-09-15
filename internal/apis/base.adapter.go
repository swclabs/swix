package apis

import (
	"swclabs/swix/internal/apis/server"
	"swclabs/swix/internal/types"
)

// NewBaseAdapter returns a new adapter wrapping around the given server
func NewBaseAdapter(server server.IServer) types.IAdapter {
	adapter := &_Adapter{
		server: server,
	}
	return adapter
}
