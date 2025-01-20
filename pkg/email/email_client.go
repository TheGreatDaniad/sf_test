package email

import (
	"fmt"
	"log"
	"net/smtp"
)

type EmailClient struct {
	SMTPHost    string
	SMTPPort    int
	Username    string
	Password    string
	SenderEmail string
}

// NewEmailClient initializes a new email client with SMTP configuration.
func NewEmailClient(host string, port int, username, password, senderEmail string) *EmailClient {
	return &EmailClient{
		SMTPHost:    host,
		SMTPPort:    port,
		Username:    username,
		Password:    password,
		SenderEmail: senderEmail,
	}
}

// SendEmail sends an email to the specified recipient.
func (e *EmailClient) SendEmail(recipient, subject, body string) error {
	auth := smtp.PlainAuth("", e.Username, e.Password, e.SMTPHost)
	address := fmt.Sprintf("%s:%d", e.SMTPHost, e.SMTPPort)

	message := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", recipient, subject, body))

	err := smtp.SendMail(address, auth, e.SenderEmail, []string{recipient}, message)
	if err != nil {
		log.Printf("Failed to send email to %s: %v", recipient, err)
		return err
	}

	log.Printf("Email sent successfully to %s", recipient)
	return nil
}
