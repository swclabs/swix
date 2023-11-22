package test

import (
	"testing"

	"swclabs/swiftcart/pkg/x/jwt"
)

func TestHash(t *testing.T) {
	pass, _ := jwt.GenPassword("12345")
	if err := jwt.ComparePassword(pass, "12345"); err != nil {
		t.Fatalf("ERROR: %s", err.Error())
	}
}
