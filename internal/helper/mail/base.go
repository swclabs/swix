package mail

import (
	"context"
	"swclabs/swipecore/pkg/tools/mailer"

	"swclabs/swipecore/internal/config"
	"swclabs/swipecore/pkg/web/components"

	"github.com/a-h/templ"
)

type Mailer struct {
	*mailer.Mailer
}

func New() *Mailer {
	return &Mailer{
		mailer.NewMailer(config.Email, config.EmailAppPassword),
	}
}

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
