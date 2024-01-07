package test

import (
	"swclabs/swiftcart/pkg/mailers"
	"swclabs/swiftcart/pkg/template"
	"testing"
)

func TestEmail(t *testing.T) {
	mailers.SendSampleHTML(template.ConfirmOrder, "iduchungho@gmail.com")
}
