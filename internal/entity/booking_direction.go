package entity

import (
	"time"
)

// BookingDirection represents a booking direction record.
type BookingDirection struct {
	BdID      int       `json:"bd_id"`
	BookingID int       `json:"booking_id"`
	Direction string    `json:"direction"`
	FlightID  int       `json:"flight_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
