package test

import (
	"testing"

	"github.com/swclabs/swipe-api/pkg/mailers"
	"github.com/swclabs/swipe-api/pkg/template"
)

func TestEmail(t *testing.T) {
	mailers.SendSampleHTML(template.ConfirmOrder, "iduchungho@gmail.com")
}
