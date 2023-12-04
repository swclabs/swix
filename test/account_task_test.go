package test

import (
	"fmt"
	"swclabs/swiftcart/internal/tasks"
	"testing"
)

func TestAccountPath(t *testing.T) {
	mapp := tasks.Path()
	for k := range mapp {
		fmt.Printf("%s\n", k)
	}
}
