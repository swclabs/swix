package services

import (
	"context"
	"log"
	"swclabs/swipecore/internal/rpc/models/greeter"
	"time"
)

type Greeter struct {
	greeter.UnimplementedGreeterServer
}

func (g *Greeter) SayHello(context context.Context, msg *greeter.Message) (*greeter.MessageReply, error) {
	log.Printf("received message: %s at time: %v\n", msg.GetMessage(), msg.GetTimestamp())
	return &greeter.MessageReply{
		Message:   "Hello " + msg.GetMessage(),
		Timestamp: time.Now().UTC().Unix(),
	}, nil
}
