package entity

import (
	"time"
)

// SeatAvailability represents seat availability for a flight.
type SeatAvailability struct {
	FlightID  int       `json:"flight_id"`
	SeatCode  string    `json:"seat_code"`
	Status    string    `json:"status"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
