package models

import (
	"github.com/google/uuid"
	"time"
)

// Account represents an user holding a wallet in the application
type Account struct {
	Id        uuid.UUID `json:"id"`
	KeyIndex  uint32    `json:"index"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
