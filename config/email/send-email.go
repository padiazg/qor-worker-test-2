package email

import (
	"bytes"
	"log"
	netMail "net/mail"
	"text/template"

	templatedata "github.com/padiazg/qor-worker-test/model/template-data"
	"github.com/qor/mailer"
)

// type SendEmailFunc func(string, templatedata.TemplateData) error
func (s *EmailSetting) SendEmail(to, subject string, data *templatedata.TemplateData) error {
	var (
		smtpmailer   *mailer.Mailer
		textTemplate = "Hi {{.First}}, welcome to the organization!"
		textBody     string
		htmlTemplate = `<h1>Hi {{.First}}, welcome to the organization!</h1>`
		htmlBody     string
		err          error
	)

	log.Printf("Data: %v", data)

	if smtpmailer, err = s.GetMailer(); err != nil {
		return err
	}

	if textBody, err = parseTemplate(textTemplate, data); err != nil {
		return err
	}

	if htmlBody, err = parseTemplate(htmlTemplate, data); err != nil {
		return err
	}

	if err := smtpmailer.Send(
		mailer.Email{
			TO:      []netMail.Address{{Address: to}},
			From:    &netMail.Address{Address: "patricio.diaz.gimenez@gmail.com"},
			Subject: subject,
			Text:    textBody,
			HTML:    htmlBody,
			// Attachments: []mailer.Attachment{},
		},
		// mailer.Template{Name: "hello", Layout: "mailers/hello", Data: data},
	); err != nil {
		return err
	}

	return nil
}

func parseTemplate(tmpl string, data *templatedata.TemplateData) (string, error) {
	var (
		textBody bytes.Buffer
		err      error
	)

	textTemplate, err := template.New("email").Parse(tmpl)
	if err != nil {
		return "", err
	}

	textTemplate.Execute(&textBody, data)

	return textBody.String(), nil
}
