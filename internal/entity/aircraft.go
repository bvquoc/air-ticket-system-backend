package entity

import (
	"time"
)

// Aircraft represents an aircraft record.
type Aircraft struct {
	AircraftID int       `json:"aircraft_id"`
	Status     string    `json:"status"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
