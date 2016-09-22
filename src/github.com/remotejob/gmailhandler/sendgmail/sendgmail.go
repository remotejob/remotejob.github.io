package sendgmail

import (
	"log"

	"net/mail"
	"net/smtp"

	gmail "github.com/scorredoira/email"
)

func Send(glogin string, gpass string, phone string, email string, skype string) {

	msg := "New appartment client phone:" + phone + " email: " + email + " skype " + skype

	m := gmail.NewMessage("Subject", msg)

	m.From = mail.Address{
		Name:    "Alex Mazurov",
		Address: "support@mazurov.eu",
	}

	m.To = []string{"support@sinelga.com"}

	err := gmail.Send("smtp.gmail.com:587", smtp.PlainAuth("", glogin, gpass, "smtp.gmail.com"), m)
	if err != nil {
		log.Println(err)
	}
}
