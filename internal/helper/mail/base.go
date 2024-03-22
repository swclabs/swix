package mail

import (
	"context"

	"github.com/a-h/templ"
	"github.com/swclabs/swipe-api/internal/config"
	"github.com/swclabs/swipe-api/pkg/tools"
	"github.com/swclabs/swipe-api/pkg/web/components"
)

type Mailler struct {
	*tools.Mailler
}

func New() *Mailler {
	return &Mailler{
		tools.NewMailer(config.Email, config.EmailAppPassword),
	}
}

func (m *Mailler) SendPurchaseOrder(to string) error {
	html := components.PurchaseOrderIndex()
	t, err := templ.ToGoHTML(context.Background(), html)
	if err != nil {
		return err
	}
	m.Mailer.SetHeader("From", m.Email)
	m.Mailer.SetHeader("To", to)
	m.Mailer.SetHeader("Subject", "Sample Message")
	m.Mailer.SetBody("text/html", string(t))

	if err := m.Dialer.DialAndSend(m.Mailer); err != nil {
		return err
	}
	return nil
}
