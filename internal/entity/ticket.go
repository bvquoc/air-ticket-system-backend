package entity

import (
	"time"
)

// Ticket represents a ticket record.
type Ticket struct {
	TicketID    int       `json:"ticket_id"`
	PassengerID int       `json:"passenger_id"`
	FlightID    int       `json:"flight_id"`
	SeatCode    string    `json:"seat_code"`
	Status      string    `json:"status"`
	Amount      float64   `json:"amount"`
	BookingID   int       `json:"booking_id"`
	IssuedAt    time.Time `json:"issued_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
