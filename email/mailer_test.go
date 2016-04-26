package email_test

import (
	. "github.com/hyperboloide/users/email"
	"github.com/hyperboloide/users/user/mock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type testTemplateBuilder struct{}

func (t testTemplateBuilder) Text(template, lang string, data interface{}) ([]byte, error) {
	return []byte("text"), nil
}

func (t testTemplateBuilder) HTML(template, lang string, data interface{}) ([]byte, error) {
	return []byte("html"), nil
}

type testMailer struct {
	dest    string
	subject string
	text    []byte
	html    []byte
}

func (t *testMailer) Send(dest string, subject string, text []byte, html []byte) error {
	t.dest = dest
	t.subject = subject
	t.text = text
	t.html = html
	return nil
}

var _ = Describe("Mailer", func() {

	It("should test mailer", func() {
		mailer := &testMailer{}
		DefaultMailer = mailer
		DefaultTemplateBuilder = testTemplateBuilder{}

		user := mock.Generate()

		msg := &Message{
			user,
			"the subject",
			"test_template",
			nil,
		}
		Ω(msg.Send()).To(BeNil())

		Ω(mailer.dest).To(Equal(user.GetEmail()))
		Ω(mailer.subject).To(Equal("the subject"))
		Ω(mailer.text).To(Equal([]byte("text")))
		Ω(mailer.html).To(Equal([]byte("html")))
	})

})
