package test

import (
	"fmt"
	"testing"

	"github.com/swclabs/swipe-api/pkg/utils"
)

func TestRandomString(t *testing.T) {
	rand := utils.RandomString(10)
	fmt.Println(rand)
}
