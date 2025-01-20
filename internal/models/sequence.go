package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Sequence struct {
	ID                  int64      `json:"id"`
	Name                string     `json:"name" validate:"required,min=3,max=255"`
	OpenTrackingEnabled bool       `json:"openTrackingEnabled"`
	ClickTrackingEnabled bool      `json:"clickTrackingEnabled"`
	Steps               []Step     `json:"steps" validate:"dive"`
	CreatedAt           time.Time  `json:"createdAt"`
	UpdatedAt           time.Time  `json:"updatedAt"`
	DeletedAt           *time.Time `json:"deletedAt,omitempty"`
}

// Validate validates the Sequence struct.
func (s *Sequence) Validate() error {
	validate := validator.New()
	return validate.Struct(s)
}