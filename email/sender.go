package emails

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
)

// EmailData contains variables for email templates
type EmailData struct {
	Subject         string
	Name           string
	AppURL         string
	VerificationURL string
}

// EmailSender handles template rendering
type EmailSender struct {
	templates map[string]*template.Template
}

// NewEmailSender initializes the email template sender
func NewEmailSender() (*EmailSender, error) {
	templates := make(map[string]*template.Template)
	
	// Load all templates from the templates directory
	tmplFiles, err := filepath.Glob("emails/templates/*.html")
	if err != nil {
		return nil, err
	}

	baseTemplate := "emails/templates/base.html"

	for _, tmpl := range tmplFiles {
		if tmpl == baseTemplate {
			continue // Skip base template as it's included separately
		}

		name := filepath.Base(tmpl)
		t, err := template.New(name).ParseFiles(tmpl, baseTemplate)
		if err != nil {
			return nil, err
		}
		templates[name] = t
	}

	return &EmailSender{templates: templates}, nil
}

// RenderTemplate renders the named template with the given data
func (es *EmailSender) RenderTemplate(templateName string, data EmailData) (string, error) {
	tmpl, ok := es.templates[templateName]
	if !ok {
		return "", os.ErrNotExist
	}

	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "base.html", data); err != nil {
		return "", err
	}

	return buf.String(), nil
}