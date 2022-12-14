package models

import (
	"github.com/google/uuid"
)

type User struct {
	UserID uuid.UUID `json:"user_id" validate:"omitempty"`
}
