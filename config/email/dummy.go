package email

import (
	"log"

	"github.com/qor/mailer"
)

type DummySender struct {
}

var _ mailer.SenderInterface = DummySender{}

func (s DummySender) Send(email mailer.Email) error {
	log.Printf("dummy email to: '%v' subject: '%s'", email.TO, email.Subject)
	return nil
}
