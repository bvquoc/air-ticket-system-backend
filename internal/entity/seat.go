package entity

import (
	"time"
)

// Seat represents a seat record.
type Seat struct {
	SeatCode   string    `json:"seat_code"`
	AircraftID int       `json:"aircraft_id"`
	Position   string    `json:"position"`
	Class      string    `json:"class"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
