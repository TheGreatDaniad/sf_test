package db

import (
	"context"
	"database/sql"
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
	tx, err := r.db.Conn.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	query := `
        INSERT INTO sequences (name, open_tracking_enabled, click_tracking_enabled, created_at, updated_at)
        VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id
    `
	var id int64
	err = tx.QueryRowContext(ctx, query, sequence.Name, sequence.OpenTrackingEnabled, sequence.ClickTrackingEnabled).Scan(&id)
	if err != nil {
		return 0, err
	}

	// Insert steps if any exist
	if len(sequence.Steps) > 0 {
		stepsQuery := `
            INSERT INTO steps (sequence_id, subject, content, step_order, wait_days, created_at, updated_at)
            VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
        `
		for _, step := range sequence.Steps {
			_, err = tx.ExecContext(ctx, stepsQuery, id, step.Subject, step.Content, step.StepOrder, step.WaitDays)
			if err != nil {
				return 0, err
			}
		}
	}

	if err = tx.Commit(); err != nil {
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
        SELECT 
            s.id, s.name, s.open_tracking_enabled, s.click_tracking_enabled, s.created_at, s.updated_at, s.deleted_at,
            st.id, st.sequence_id, st.subject, st.content, st.step_order, st.wait_days, st.created_at, st.updated_at, st.deleted_at
        FROM sequences s
        LEFT JOIN steps st ON s.id = st.sequence_id
        WHERE s.id = $1
        ORDER BY st.step_order ASC
    `
	sequence := &models.Sequence{}
	rows, err := r.db.Conn.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var steps []models.Step
	first := true

	for rows.Next() {
		var step models.Step
		var stepID, sequenceID sql.NullInt64
		var subject, content sql.NullString
		var stepOrder, waitDays sql.NullInt32
		var stepCreatedAt, stepUpdatedAt, stepDeletedAt sql.NullTime

		err := rows.Scan(
			&sequence.ID,
			&sequence.Name,
			&sequence.OpenTrackingEnabled,
			&sequence.ClickTrackingEnabled,
			&sequence.CreatedAt,
			&sequence.UpdatedAt,
			&sequence.DeletedAt,
			&stepID,
			&sequenceID,
			&subject,
			&content,
			&stepOrder,
			&waitDays,
			&stepCreatedAt,
			&stepUpdatedAt,
			&stepDeletedAt,
		)
		if err != nil {
			return nil, err
		}

		// Only set sequence fields on first row
		if first {
			first = false
		}

		// If we have a valid step ID, add the step
		if stepID.Valid {
			step.ID = stepID.Int64
			step.SequenceID = sequenceID.Int64
			step.Subject = subject.String
			step.Content = content.String
			step.StepOrder = int(stepOrder.Int32)
			step.WaitDays = int(waitDays.Int32)
			step.CreatedAt = stepCreatedAt.Time
			step.UpdatedAt = stepUpdatedAt.Time
			if stepDeletedAt.Valid {
				step.DeletedAt = &stepDeletedAt.Time
			}
			steps = append(steps, step)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	sequence.Steps = steps
	return sequence, nil
}
