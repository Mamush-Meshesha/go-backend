package config

type EmailConfig struct {
	SMTPHost	string
	SMTPPort	int
	SMTPUsername string
	SMTPPassword string
	FromEmail	string
	UseTLS       bool 
	AppURL       string 
}

func LoadEmailConfig() *EmailConfig {
	return &EmailConfig{
		SMTPHost: 	getEnv("SMTP_HOST", ",mailhog"),
		SMTPPort: 	getEnvAsInt("SMTP_PORT", 1025),
		SMTPUsername: getEnv("SMTP_USERNAME", ""),
		SMTPPassword: getEnv("SMTP_PASSWORD", ""),
		FromEmail: getEnv("FROM_EMAIL", "nereply@golang.com"),
		UseTLS:       true,
	}
}