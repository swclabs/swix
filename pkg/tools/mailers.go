package tools

import (
	"gopkg.in/gomail.v2"
)

type Mailler struct {
	Email       string
	appPassword string
	Mailer      *gomail.Message
	Dialer      *gomail.Dialer
	Host        string
	Port        int
}

func NewMailer(email, appPassword string) *Mailler {
	return &Mailler{
		Email:       email,
		appPassword: appPassword,
		Mailer:      gomail.NewMessage(),
		Dialer:      gomail.NewDialer("smtp.gmail.com", 587, email, appPassword),
		Host:        "smtp.gmail.com",
		Port:        587,
	}
}

func (m *Mailler) Config(email, appPassword string) {
	m.Email = email
	m.appPassword = appPassword
}
