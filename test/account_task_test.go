package test

import (
	"fmt"
	"swclabs/swiftcart/internal/delivery/msg"
	"testing"
)

func TestAccountPath(t *testing.T) {
	mapp := msg.Path()
	for k := range mapp {
		fmt.Printf("%s\n", k)
	}
}
