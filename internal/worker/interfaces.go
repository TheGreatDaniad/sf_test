package worker

import (
	"context"
	"sf_test/internal/models"
)

type EmailScheduler interface {
	ScheduleEmails(ctx context.Context, sequenceID int64) error
}

type EmailSender interface {
	SendEmail(ctx context.Context, step models.Step, recipient string) error
}
