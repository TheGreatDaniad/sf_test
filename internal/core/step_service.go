package core

import (
	"context"
	"sf_test/internal/db"
	"sf_test/internal/models"
)

type stepService struct {
	repo db.StepRepository
}

func NewStepService(repo db.StepRepository) StepService {
	return &stepService{repo: repo}
}

func (s *stepService) CreateStep(ctx context.Context, step *models.Step) (int64, error) {
	// Validate the step model
	if err := step.Validate(); err != nil {
		return 0, err
	}

	// Save the step to the repository
	return s.repo.Create(ctx, step)
}

func (s *stepService) UpdateStep(ctx context.Context, step *models.Step) error {
	// Validate the step model
	if err := step.Validate(); err != nil {
		return err
	}

	// Update the step in the repository
	return s.repo.Update(ctx, step)
}

func (s *stepService) DeleteStep(ctx context.Context, id int64) error {
	// Delete the step from the repository
	return s.repo.Delete(ctx, id)
}

func (s *stepService) ListSteps(ctx context.Context, sequenceID int64) ([]*models.Step, error) {
	// Retrieve steps for the given sequence ID
	steps, err := s.repo.ListBySequenceID(ctx, sequenceID)
	if err != nil {
		return nil, err
	}
	return steps, nil
}
