package email

import (
	"fmt"

	"github.com/qor/mailer"
	"github.com/qor/mailer/gomailer"
	"gopkg.in/gomail.v2"
)

type EmailSetting struct {
	Host     string
	Username string
	Password string
	Port     int
}

func (s *EmailSetting) GetMailer() (*mailer.Mailer, error) {
	var (
		senderCloser gomail.SendCloser
		sender       mailer.SenderInterface
		dialer       *gomail.Dialer
		err          error
	)

	if s.Host == "" {
		sender = DummySender{}
	} else {
		dialer = gomail.NewDialer(
			s.Host,
			s.Port,
			s.Username,
			s.Password,
		)

		if senderCloser, err = dialer.Dial(); err != nil {
			return nil, fmt.Errorf("failed to init mailer %v", err)
		}
		sender = gomailer.New(&gomailer.Config{Sender: senderCloser})

	}

	// configure mailer
	config := &mailer.Config{
		Sender: sender,
	}

	return mailer.New(config), nil
}

func (s *EmailSetting) Live() (*mailer.Mailer, error) {
	var (
		senderCloser gomail.SendCloser
		err          error
	)

	dialer := gomail.NewDialer(
		s.Host,
		s.Port,
		s.Username,
		s.Password,
	)

	if senderCloser, err = dialer.Dial(); err != nil {
		return nil, fmt.Errorf("failed to init mailer %v", err)
	}

	return mailer.New(&mailer.Config{
		Sender: gomailer.New(&gomailer.Config{Sender: senderCloser}),
	}), nil

}

func (s *EmailSetting) Dummy() (*DummySender, error) {
	return &DummySender{}, nil
}
