package db

import (
	"context"
	"errors"
	"sf_test/internal/models"
)

type StepRepository interface {
	Create(ctx context.Context, step *models.Step) (int64, error)
	Update(ctx context.Context, step *models.Step) error
	Delete(ctx context.Context, id int64) error
	ListBySequenceID(ctx context.Context, sequenceID int64) ([]*models.Step, error)
}

type stepRepo struct {
	db *DB
}

func NewStepRepository(db *DB) StepRepository {
	return &stepRepo{db: db}
}

func (r *stepRepo) Create(ctx context.Context, step *models.Step) (int64, error) {
	query := `
        INSERT INTO steps (sequence_id, subject, content, step_order, wait_days, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) RETURNING id
    `
	var id int64
	err := r.db.Conn.QueryRowContext(ctx, query, step.SequenceID, step.Subject, step.Content, step.StepOrder, step.WaitDays).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *stepRepo) Update(ctx context.Context, step *models.Step) error {
	query := `
        UPDATE steps
        SET subject = $1, content = $2, updated_at = NOW()
        WHERE id = $3
    `
	result, err := r.db.Conn.ExecContext(ctx, query, step.Subject, step.Content, step.ID)
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

func (r *stepRepo) Delete(ctx context.Context, id int64) error {
	query := `
        DELETE FROM steps
        WHERE id = $1
    `
	result, err := r.db.Conn.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows deleted")
	}
	return nil
}

func (r *stepRepo) ListBySequenceID(ctx context.Context, sequenceID int64) ([]*models.Step, error) {
	query := `
        SELECT id, sequence_id, subject, content, step_order, wait_days, created_at, updated_at, deleted_at
        FROM steps
        WHERE sequence_id = $1
        ORDER BY step_order
    `
	rows, err := r.db.Conn.QueryContext(ctx, query, sequenceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var steps []*models.Step
	for rows.Next() {
		step := &models.Step{}
		if err := rows.Scan(&step.ID, &step.SequenceID, &step.Subject, &step.Content, &step.StepOrder, &step.WaitDays, &step.CreatedAt, &step.UpdatedAt, &step.DeletedAt); err != nil {
			return nil, err
		}
		steps = append(steps, step)
	}
	return steps, nil
}
