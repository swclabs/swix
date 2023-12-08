package test

import (
	"swclabs/swiftcart/pkg/mailers"
	"swclabs/swiftcart/pkg/template"
	"testing"
)

func TestEmail(t *testing.T) {
	mailers.SendHTML(template.ConfirmOrder, "iduchungho@gmail.com")
}
