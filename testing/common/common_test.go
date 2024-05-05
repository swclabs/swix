package test

import (
	"fmt"
	"maps"
	"swclabs/swipecore/pkg/tools/jwt"
	"testing"

	"swclabs/swipecore/pkg/utils"
)

func TestHash(t *testing.T) {
	pass, _ := jwt.GenPassword("12345")
	if err := jwt.ComparePassword(pass, "12345"); err != nil {
		t.Fatalf("ERROR: %s", err.Error())
	}
}

func TestWorkerPath(t *testing.T) {
	map1 := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	map2 := map[string]int{
		"4": 4,
		"5": 5,
		"6": 6,
	}
	maps.Copy(map1, map2)
}

func TestRandomString(t *testing.T) {
	rand := utils.RandomString(10)
	if len(rand) != 10 {
		t.Error("should have 10 random strings")
	}
}

func TestValidEmail(t *testing.T) {
	for i, email := range []string{
		"good@exmaple.com",
		"bad-example",
	} {
		isEmail := utils.IsEmail(email)
		if i == 0 && !isEmail {
			t.Error("should have valid email: " + email)
		}
		if i == 1 && isEmail {
			t.Error("should have invalid email: " + email)
		}
	}
}

func TestStmt(t *testing.T) {
	queryHandler := func(sql string, args ...interface{}) {
		fmt.Print(sql)
		fmt.Println(args...)
	}
	queryHandler("select * from", "table", "row", "column")
}
