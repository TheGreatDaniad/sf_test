package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Step struct {
	ID         int64      `json:"id"`
	SequenceID int64      `json:"sequenceId"`
	Subject    string     `json:"subject" validate:"required,min=3,max=255"`
	Content    string     `json:"content" validate:"required"`
	StepOrder  int        `json:"stepOrder" validate:"gte=0"`
	WaitDays   int        `json:"waitDays" validate:"gte=0"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt,omitempty"`
}

// Validate validates the Step struct.
func (s *Step) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}
