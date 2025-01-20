package worker

import (
	"context"
	"errors"
	"log"

	"sf_test/internal/models"
	"sf_test/pkg/email"
)

type emailWorker struct {
	emailClient *email.EmailClient
}

// NewEmailWorker initializes the email worker with an email client.
func NewEmailWorker(client *email.EmailClient) EmailSender {
	return &emailWorker{emailClient: client}
}

// SendEmail sends an email to a recipient for the given step using the email client.
func (w *emailWorker) SendEmail(ctx context.Context, step models.Step, recipient string) error {
	if step.Subject == "" || step.Content == "" {
		return errors.New("invalid step: missing subject or content")
	}

	log.Printf("Preparing to send email to %s with Subject: %s", recipient, step.Subject)

	// Send the email using the email client
	err := w.emailClient.SendEmail(recipient, step.Subject, step.Content)
	if err != nil {
		log.Printf("Failed to send email to %s: %v", recipient, err)
		return err
	}

	log.Printf("Email sent successfully to %s", recipient)
	return nil
}
