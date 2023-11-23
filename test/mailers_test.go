package test

import (
	"swclabs/swiftcart/pkg/template"
	"swclabs/swiftcart/pkg/x/mailers"
	"testing"
)

func TestEmail(t *testing.T) {
	mailers.SendHTML(template.ConfirmOrder, "iduchungho@gmail.com")
}
