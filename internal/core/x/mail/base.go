// Package mail connect to mail server
package mail

import (
	"context"
	"swclabs/swipex/internal/config"
	"swclabs/swipex/pkg/lib/mailer"

	"swclabs/swipex/pkg/components"

	"github.com/a-h/templ"
)

// Mailer struct to send email
type Mailer struct {
	*mailer.Mailer
}

// New returns a new Mailer object
func New() *Mailer {
	return &Mailer{
		mailer.NewMailer(config.Email, config.EmailAppPassword),
	}
}

// SendPurchaseOrder sends a purchase order email
func (m *Mailer) SendPurchaseOrder(to string) error {
	html := components.PurchaseOrderIndex()
	t, err := templ.ToGoHTML(context.Background(), html)
	if err != nil {
		return err
	}

	m.Message.SetHeader("From", m.Email)
	m.Message.SetHeader("To", to)
	m.Message.SetHeader("Subject", "Sample Message")
	m.Message.SetBody("text/html", string(t))

	if err := m.Dialer.DialAndSend(m.Message); err != nil {
		return err
	}
	return nil
}
