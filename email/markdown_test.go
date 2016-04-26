package email_test

import (
	html "html/template"
	text "text/template"

	. "github.com/hyperboloide/users/email"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type testProvider struct{}

func (p testProvider) Get(templateName, lang string) (*text.Template, error) {
	return text.New("test").Parse("hello **{{ .Name }}**.")
}

func (p testProvider) Body(lang string) (*html.Template, error) {
	return html.New("test").Parse("<html>{{ . }}</html>")
}

var _ = Describe("Markdown", func() {

	It("should test the Provider", func() {
		provider := testProvider{}

		builder := NewMarkdownBuilder(provider)
		Ω(builder).ToNot(BeNil())

		buff, err := builder.Text("test", "en", struct {
			Name string
		}{"User"})
		Ω(err).To(BeNil())
		Ω(buff).ToNot(BeNil())
		Ω(string(buff[:])).To(Equal(`hello **User**.`))

		buff, err = builder.HTML("test", "en", struct {
			Name string
		}{"User"})
		Ω(err).To(BeNil())
		Ω(buff).ToNot(BeNil())
		Ω(string(buff[:])).To(Equal("<html><p>hello <strong>User</strong>.</p>\n</html>"))
	})
})
