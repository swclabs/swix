package tools

import (
	"gopkg.in/gomail.v2"
)

type Mailer struct {
	Email       string
	appPassword string
	Message     *gomail.Message
	Dialer      *gomail.Dialer
	Host        string
	Port        int
}

func NewMailer(email, appPassword string) *Mailer {
	return &Mailer{
		Email:       email,
		appPassword: appPassword,
		Message:     gomail.NewMessage(),
		Dialer:      gomail.NewDialer("smtp.gmail.com", 587, email, appPassword),
		Host:        "smtp.gmail.com",
		Port:        587,
	}
}

func (m *Mailer) Config(email, appPassword string) {
	m.Email = email
	m.appPassword = appPassword
}
