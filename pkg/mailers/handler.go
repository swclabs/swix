package mailers

import (
	"bytes"
	"errors"
	"log"
	"text/template"
)

func SendSampleHTML(tempPath, to string) error {
	if mailer == nil {
		return errors.New("missing app passsword & email address specified")
	}
	
	t, err := template.ParseFiles(tempPath)
	if err != nil {
		log.Fatal(err)
	}
	
	var body bytes.Buffer
	err = t.Execute(&body, struct{ Name string }{Name: "Kawasaki"})
	if err != nil {
		return err
	}

	lock.Lock()
	defer lock.Unlock()

	mailer.SetHeader("From", Env.Email)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", "Sample Message")
	mailer.SetBody("text/html", body.String())
	// mail.SetAddressHeader("Cc", "dan@example.com", "Dan")
	// mail.Attach("/home/Alex/lolcat.jpg")

	// Send the email to Bob, Cora and Dan.
	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}
	return nil
}
