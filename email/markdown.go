package email

import (
	"bytes"
	"fmt"
	html "html/template"
	text "text/template"

	"github.com/russross/blackfriday"
)

// MarkdownProvider provides templates for markdown Builder.
type MarkdownProvider interface {
	// Get the markdown template by name and language (returns *text/template.Template).
	Get(templateName, lang string) (*text.Template, error)
	// Body for html messages by language (returns *html/template.Template).
	Body(lang string) (*html.Template, error)
}

// MarkdownBuilder generates email templates from markdown.
type MarkdownBuilder struct {
	Provider MarkdownProvider
}

// NewMarkdownBuilder creates a new MarkdownBuilder from a Provider.
func NewMarkdownBuilder(provider MarkdownProvider) *MarkdownBuilder {
	return &MarkdownBuilder{provider}
}

// Text generates an email text body.
func (b *MarkdownBuilder) Text(template, lang string, data interface{}) ([]byte, error) {
	buff := &bytes.Buffer{}
	if tmpl, err := b.Provider.Get(template, lang); err != nil {
		return nil, err
	} else if err := tmpl.Execute(buff, data); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

// HTML generates an HTML body.
func (b *MarkdownBuilder) HTML(template, lang string, data interface{}) ([]byte, error) {
	tmpl, err := b.Provider.Body(lang)
	if err != nil {
		return nil, err
	} else if tmpl == nil {
		return nil, fmt.Errorf("Body html template for language '%s' not found", lang)
	}

	txt, err := b.Text(template, lang, data)
	if err != nil {
		return nil, err
	}

	buff := &bytes.Buffer{}
	content := blackfriday.MarkdownCommon(txt)
	if err := tmpl.Execute(buff, html.HTML(content)); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
