package entity

import (
	"time"
)

// Airport represents an airport record.
type Airport struct {
	AirportCode string    `json:"airport_code"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
