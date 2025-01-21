package models

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type Sequence struct {
	ID                   int64      `json:"id"`
	Name                 string     `json:"name" validate:"required,min=3,max=255"`
	OpenTrackingEnabled  bool       `json:"openTrackingEnabled"`
	ClickTrackingEnabled bool       `json:"clickTrackingEnabled"`
	Steps                []Step     `json:"steps" validate:"dive"`
	CreatedAt            time.Time  `json:"createdAt"`
	UpdatedAt            time.Time  `json:"updatedAt"`
	DeletedAt            *time.Time `json:"deletedAt,omitempty"`
}

// Validate validates the Sequence struct.
func (s *Sequence) Validate() error {
	validate := validator.New()

	// Validate the struct using the default rules.
	if err := validate.Struct(s); err != nil {
		return err
	}

	// Custom validation for unique stepOrder in Steps.
	stepOrderMap := make(map[int]bool)
	for _, step := range s.Steps {
		if stepOrderMap[step.StepOrder] {
			return errors.New("stepOrder values must be unique in Steps")
		}
		stepOrderMap[step.StepOrder] = true
	}

	return nil
}
