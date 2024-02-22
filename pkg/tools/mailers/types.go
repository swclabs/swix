package mailers

import (
	"sync"

	"gopkg.in/gomail.v2"
)

const Host = "smtp.gmail.com"
const Port = 587

var mailer *gomail.Message = nil
var dialer *gomail.Dialer = nil
var lockInit *sync.Mutex = &sync.Mutex{}
var lock *sync.Mutex = &sync.Mutex{}

func init() {
	if Env.AppPassword != "" && Env.Email != "" {
		if mailer == nil {
			lockInit.Lock()
			defer lockInit.Unlock()
			if mailer == nil {
				mailer = gomail.NewMessage()
				dialer = gomail.NewDialer(Host, Port, Env.Email, Env.AppPassword)
			}
		}
		return
	}
}

type MailerEnv struct {
	Email       string `json:"email"`
	AppPassword string `json:"app_password"`
}
