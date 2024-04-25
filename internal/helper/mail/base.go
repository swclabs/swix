package mail

import (
	"context"

	"github.com/a-h/templ"
	"swclabs/swipe-api/internal/config"
	"swclabs/swipe-api/pkg/tools"
	"swclabs/swipe-api/pkg/web/components"
)

type Mailer struct {
	*tools.Mailer
}

func New() *Mailer {
	return &Mailer{
		tools.NewMailer(config.Email, config.EmailAppPassword),
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
