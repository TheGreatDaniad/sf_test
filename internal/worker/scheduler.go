package worker

import (
	"context"
	"log"
	"time"

	"sf_test/internal/core"
)

type scheduler struct {
	stepService core.StepService
	emailSender EmailSender
}

// NewScheduler initializes the scheduler with step service and email sender.
func NewScheduler(stepService core.StepService, emailSender EmailSender) EmailScheduler {
	return &scheduler{stepService: stepService, emailSender: emailSender}
}

// ScheduleEmails schedules emails for the given sequence.
func (s *scheduler) ScheduleEmails(ctx context.Context, sequenceID int64) error {
	// Fetch steps for the given sequence ID
	steps, err := s.stepService.ListSteps(ctx, sequenceID)
	if err != nil {
		log.Printf("Failed to fetch steps for sequence %d: %v", sequenceID, err)
		return err
	}

	// Iterate through each step and schedule emails
	for _, step := range steps {
		log.Printf("Scheduling emails for Step %d (Order %d) in Sequence %d", step.ID, step.StepOrder, sequenceID)

		// Simulate recipients for the step (this should be replaced with actual recipients logic)
		recipients := []string{"recipient1@example.com", "recipient2@example.com"}

		// Send emails to all recipients
		for _, recipient := range recipients {
			err := s.emailSender.SendEmail(ctx, *step, recipient)
			if err != nil {
				log.Printf("Failed to send email to %s for Step %d: %v", recipient, step.ID, err)
				continue
			}
			log.Printf("Successfully sent email to %s for Step %d", recipient, step.ID)
		}

		// Wait for the defined number of days before the next step
		log.Printf("Waiting for %d days before scheduling the next step...", step.WaitDays)
		time.Sleep(time.Duration(step.WaitDays) * 24 * time.Hour) // Simulating wait days
	}

	log.Printf("Finished scheduling emails for Sequence %d", sequenceID)
	return nil
}
