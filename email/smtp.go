package email

import (
	"errors"

	"github.com/go-gomail/gomail"
)

// SMTP can send email through the SMTP protocol
type SMTP struct {
	Host     string
	Port     int
	User     string
	Password string
	Sender   string
}

// Send an email
func (s SMTP) Send(dest string, subject string, text []byte, html []byte) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", s.Sender)
	msg.SetHeader("To", dest)
	msg.SetHeader("Subject", subject)

	if text == nil && html == nil {
		return errors.New("Email needs at least a text or an html body")
	} else if text == nil {
		msg.SetBody("text/html", string(html[:]))
	} else if html == nil {
		msg.SetBody("text/plain", string(text[:]))
	} else {
		msg.SetBody("text/plain", string(text[:]))
		msg.AddAlternative("text/html", string(html[:]))
	}

	return gomail.NewDialer(
		s.Host,
		s.Port,
		s.User,
		s.Password).DialAndSend(msg)
}
