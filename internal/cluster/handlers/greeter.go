// Package service defines all serivce in base package.
package handlers

import (
	"context"
	"fmt"
	"swclabs/swix/internal/cluster/proto/greeter"
	"time"
)

// Greeter is the service for Greeter.
type Greeter struct {
	greeter.UnimplementedGreeterServer
}

// NewGreeter creates a new Greeter service.
func NewGreeter() greeter.GreeterServer {
	return &Greeter{}
}

// SayHello implements GreeterServer.
func (g *Greeter) SayHello(_ context.Context, msg *greeter.MessageRequest) (*greeter.MessageReply, error) {
	body := msg.GetBody()
	if body == nil || body.Msg == nil {
		return nil, fmt.Errorf("body is required")
	}
	return &greeter.MessageReply{
		Message:   "Reply " + body.Msg.GetValue(),
		Timestamp: time.Now().UTC().Unix(),
	}, nil
}

// HeathCheck implements GreeterServer.
func (g *Greeter) HeathCheck(_ context.Context, _ *greeter.StringMessage) (*greeter.StringMessage, error) {
	return &greeter.StringMessage{
		Value: "Ok",
	}, nil
}
