package services

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"todo/config"
	"todo/email"
)

type EmailService struct {
	cfg    *config.EmailConfig
	sender *emails.EmailSender
}

func NewEmailService(cfg *config.EmailConfig) (*EmailService, error) {
	sender, err := emails.NewEmailSender()
	if err != nil {
		return nil, fmt.Errorf("failed to create email sender: %w", err)
	}
	return &EmailService{
		cfg:    cfg,
		sender: sender,
	}, nil
}

func (s *EmailService) SendWelcomeEmail(to, name string) error {
	htmlContent, err := s.sender.RenderTemplate("welcome.html", emails.EmailData{
		Subject: "Welcome to Todo App!",
		Name:    name,
		AppURL:  s.cfg.AppURL,
	})
	if err != nil {
		return fmt.Errorf("failed to render welcome template: %w", err)
	}

	return s.sendEmail(to, "Welcome to Todo App", htmlContent)
}

func (s *EmailService) SendVerificationEmail(to, name, token string) error {
	verifyURL := fmt.Sprintf("%s/auth/verify?token=%s", s.cfg.AppURL, token)
	
	htmlContent, err := s.sender.RenderTemplate("verification.html", emails.EmailData{
		Subject:         "Verify Your Email",
		Name:           name,
		VerificationURL: verifyURL,
	})
	if err != nil {
		return fmt.Errorf("failed to render verification template: %w", err)
	}

	return s.sendEmail(to, "Verify Your Email", htmlContent)
}

func (s *EmailService) sendEmail(to, subject, body string) error {
	// Construct MIME message
	headers := map[string]string{
		"From":         s.cfg.FromEmail,
		"To":           to,
		"Subject":      subject,
		"MIME-Version": "1.0",
		"Content-Type": "text/html; charset=UTF-8",
	}

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Set up authentication
	auth := smtp.PlainAuth("", s.cfg.SMTPUsername, s.cfg.SMTPPassword, s.cfg.SMTPHost)

	// Connect to server
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", s.cfg.SMTPHost, s.cfg.SMTPPort), &tls.Config{
		ServerName: s.cfg.SMTPHost,
	})
	if err != nil {
		return fmt.Errorf("TLS connection failed: %w", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, s.cfg.SMTPHost)
	if err != nil {
		return fmt.Errorf("SMTP client creation failed: %w", err)
	}
	defer client.Close()

	// Authenticate
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("authentication failed: %w", err)
	}

	// Set sender and recipient
	if err = client.Mail(s.cfg.FromEmail); err != nil {
		return fmt.Errorf("sender setup failed: %w", err)
	}
	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("recipient setup failed: %w", err)
	}

	// Send email body
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("data command failed: %w", err)
	}
	
	_, err = w.Write([]byte(message))
	if err != nil {
		w.Close()
		return fmt.Errorf("message writing failed: %w", err)
	}
	
	if err = w.Close(); err != nil {
		return fmt.Errorf("message closing failed: %w", err)
	}

	return client.Quit()
}