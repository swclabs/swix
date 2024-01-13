package test

import (
	"fmt"
	"swclabs/swiftcart/internal/messaging"
	"testing"
)

func TestAccountPath(t *testing.T) {
	mapp := messaging.Controller()
	for k := range mapp {
		fmt.Printf("%s\n", k)
	}
}
