// Package mailer implements mailer
package mailer

import (
	"gopkg.in/gomail.v2"
)

// Mailer struct define the Mailer object
type Mailer struct {
	Email       string
	appPassword string
	Message     *gomail.Message
	Dialer      *gomail.Dialer
	Host        string
	Port        int
}

// NewMailer creates a new Mailer object
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

// Config set email and app password
func (m *Mailer) Config(email, appPassword string) {
	m.Email = email
	m.appPassword = appPassword
}
