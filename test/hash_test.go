package test

import (
	"example/swiftcart/pkg/utils"
	"testing"
)

func TestHash(t *testing.T) {
	pass, _ := utils.GenPassword("12345")
	if err := utils.ComparePassword(pass, "12345"); err != nil {
		t.Fatalf("ERROR: %s", err.Error())
	}
}
