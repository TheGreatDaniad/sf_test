package core

import (
	"context"
	"database/sql"
	"errors"
	"sf_test/internal/db"
	"sf_test/internal/models"
)

type sequenceService struct {
	repo db.SequenceRepository
}

func NewSequenceService(repo db.SequenceRepository) SequenceService {
	return &sequenceService{repo: repo}
}

func (s *sequenceService) CreateSequence(ctx context.Context, sequence *models.Sequence) (int64, error) {
	// Validate the sequence model
	if err := sequence.Validate(); err != nil {
		return 0, err
	}

	// Save the sequence to the repository
	return s.repo.Create(ctx, sequence)
}

func (s *sequenceService) UpdateTracking(ctx context.Context, id int64, openTracking, clickTracking bool) error {
	// Update tracking flags in the repository
	return s.repo.UpdateTracking(ctx, id, openTracking, clickTracking)
}

func (s *sequenceService) GetSequence(ctx context.Context, id int64) (*models.Sequence, error) {
	// Retrieve the sequence from the repository
	sequence, err := s.repo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("sequence not found")
		}
		return nil, err
	}
	return sequence, nil
}
