package email

import (
	netMail "net/mail"

	templatedata "github.com/padiazg/qor-worker-test/model/template-data"
	"github.com/qor/mailer"
)

// type SendEmailFunc func(string, templatedata.TemplateData) error

func (s *EmailSetting) SendEmail(to, subject string, data templatedata.TemplateData) error {
	var (
		smtpmailer *mailer.Mailer
		err        error
		// text = template.New("email")
		// textTemplate = `Hi {{.Name}}, welcome to the organization!`
		// htmlTemplate = `<h1>Hi {{.Name}}, welcome to the organization!</h1>`
	)

	// text.Parse(textTemplate)

	if smtpmailer, err = s.GetMailer(); err != nil {
		return err
	}

	if err := smtpmailer.Send(
		mailer.Email{
			TO:      []netMail.Address{{Address: to}},
			From:    &netMail.Address{Address: "patricio.diaz.gimenez@gmail.com"},
			Subject: subject,
			// Text:    "Hi {{.Name}}, welcome to the organization!",
			// HTML:    "<h1>Hi {{.Name}}, welcome to the organization!</h1>",
			// Attachments: []mailer.Attachment{},
		},
		mailer.Template{Name: "hello", Layout: "mailers/hello", Data: data},
	); err != nil {
		return err
	}

	return nil
}
