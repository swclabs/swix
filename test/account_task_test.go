package test

import (
	"fmt"
	"swclabs/swiftcart/delivery/messaging"
	"testing"
)

func TestAccountPath(t *testing.T) {
	mapp := messaging.Path()
	for k := range mapp {
		fmt.Printf("%s\n", k)
	}
}
