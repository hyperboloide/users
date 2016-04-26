package email

import (
	"errors"

	"github.com/hyperboloide/users/user"
)

// Mailer sends Messages.
type Mailer interface {
	// Send a message.
	Send(dest string, subject string, text []byte, html []byte) error
}

// TemplateBuilder generates the content of a Message.
type TemplateBuilder interface {
	// GenerateText for data in language lang
	Text(template, lang string, data interface{}) ([]byte, error)
	// GenerateHTML for data in language lang
	HTML(template, lang string, data interface{}) ([]byte, error)
}

var (
	// DefaultMailer is the mailer used.
	DefaultMailer Mailer
	// DefaultTemplateBuilder is used to generate templates
	DefaultTemplateBuilder TemplateBuilder
)

// Message represents an email to be sent.
type Message struct {
	Dest     user.User   `json:"dest"`
	Subject  string      `json:"subject"`
	Template string      `json:"template"`
	Data     interface{} `json:"data"`
}

// Send a Message
func (msg Message) Send() error {
	if msg.Dest.GetEmail() == "" {
		return nil
	} else if DefaultMailer == nil {
		return errors.New("DefaultMailer not set")
	} else if DefaultTemplateBuilder == nil {
		return errors.New("DefaultTemplateBuilder not set")
	}

	text, err := DefaultTemplateBuilder.Text(
		msg.Template,
		msg.Dest.GetLanguage(),
		msg.Data)
	if err != nil {
		return err
	}

	html, err := DefaultTemplateBuilder.HTML(
		msg.Template,
		msg.Dest.GetLanguage(),
		msg.Data)
	if err != nil {
		return err
	}

	return DefaultMailer.Send(msg.Dest.GetEmail(), msg.Subject, text, html)
}
