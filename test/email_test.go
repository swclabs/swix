package test

import (
	"testing"

	"swclabs/swipe-api/internal/helper/mail"
)

func TestEmail(t *testing.T) {
	m := mail.New()
	err := m.SendPurchaseOrder("iduchungho@gmail.com")
	if err != nil {
		t.Fatal(err)
	}
}
