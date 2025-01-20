package db

import (
	"context"
	"errors"
	"sf_test/internal/models"
)

type SequenceRepository interface {
	Create(ctx context.Context, sequence *models.Sequence) (int64, error)
	UpdateTracking(ctx context.Context, id int64, openTracking, clickTracking bool) error
	Get(ctx context.Context, id int64) (*models.Sequence, error)
}

type sequenceRepo struct {
	db *DB
}

func NewSequenceRepository(db *DB) SequenceRepository {
	return &sequenceRepo{db: db}
}

func (r *sequenceRepo) Create(ctx context.Context, sequence *models.Sequence) (int64, error) {
	query := `
        INSERT INTO sequences (name, open_tracking_enabled, click_tracking_enabled, created_at, updated_at)
        VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id
    `
	var id int64
	err := r.db.Conn.QueryRowContext(ctx, query, sequence.Name, sequence.OpenTrackingEnabled, sequence.ClickTrackingEnabled).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *sequenceRepo) UpdateTracking(ctx context.Context, id int64, openTracking, clickTracking bool) error {
	query := `
        UPDATE sequences
        SET open_tracking_enabled = $1, click_tracking_enabled = $2, updated_at = NOW()
        WHERE id = $3
    `
	result, err := r.db.Conn.ExecContext(ctx, query, openTracking, clickTracking, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows updated")
	}
	return nil
}

func (r *sequenceRepo) Get(ctx context.Context, id int64) (*models.Sequence, error) {
	query := `
        SELECT id, name, open_tracking_enabled, click_tracking_enabled, created_at, updated_at, deleted_at
        FROM sequences
        WHERE id = $1
    `
	sequence := &models.Sequence{}
	err := r.db.Conn.QueryRowContext(ctx, query, id).
		Scan(&sequence.ID, &sequence.Name, &sequence.OpenTrackingEnabled, &sequence.ClickTrackingEnabled, &sequence.CreatedAt, &sequence.UpdatedAt, &sequence.DeletedAt)
	if err != nil {
		return nil, err
	}
	return sequence, nil
}
