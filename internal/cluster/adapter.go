package cluster

import (
	"context"
	"swclabs/swix/internal/cluster/port"
	"swclabs/swix/internal/types"
)

type Adapter struct {
	gateway IGateway
	cluster ICluster
}

func NewAdapter(cluster ICluster, gate IGateway) types.IAdapter {
	return &Adapter{
		gateway: gate,
		cluster: cluster,
	}
}

// Run implements types.IAdapter.
func (a *Adapter) Run(prop string) error {
	a.cluster.ServeNode(port.Greeter)
	// connect the gateway to the cluster
	if err := a.gateway.Connect(context.Background(), a.cluster); err != nil {
		return err
	}
	return a.gateway.ListenAndServe("localhost:8080")
}
