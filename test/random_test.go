package test

import (
	"fmt"
	"swclabs/swiftcart/pkg/utils"
	"testing"
)

func TestRandomString(t *testing.T) {
	rand := utils.RandomString(10)
	fmt.Println(rand)
}
