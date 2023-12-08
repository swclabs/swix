package test

import (
	"fmt"
	"maps"
	"testing"

	"swclabs/swiftcart/pkg/jwt"
)

func TestHash(t *testing.T) {
	pass, _ := jwt.GenPassword("12345")
	if err := jwt.ComparePassword(pass, "12345"); err != nil {
		t.Fatalf("ERROR: %s", err.Error())
	}
}

func TestWorkerPath(t *testing.T) {
	map1 := map[string]int{
		// "1": 1,
		// "2": 2,
		// "3": 3,
	}
	map2 := map[string]int{
		"4": 4,
		"5": 5,
		"6": 6,
	}
	maps.Copy(map1, map2)
	for k, v := range map1 {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}
}
