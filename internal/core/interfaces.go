package core

import (
	"context"
	"sf_test/internal/models"
)

// SequenceService defines the interface for sequence-related operations.
type SequenceService interface {
	CreateSequence(ctx context.Context, sequence *models.Sequence) (int64, error)
	UpdateTracking(ctx context.Context, id int64, openTracking, clickTracking bool) error
	GetSequence(ctx context.Context, id int64) (*models.Sequence, error)
}

// StepService defines the interface for step-related operations.
type StepService interface {
	CreateStep(ctx context.Context, step *models.Step) (int64, error)
	UpdateStep(ctx context.Context, step *models.Step) error
	DeleteStep(ctx context.Context, id int64) error
	ListSteps(ctx context.Context, sequenceID int64) ([]*models.Step, error)
}
