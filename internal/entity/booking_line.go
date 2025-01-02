package entity

import (
	"time"
)

// BookingLine represents a line in a booking.
type BookingLine struct {
	BlID           int       `json:"bl_id"`
	BdID           int       `json:"bd_id"`
	Type           string    `json:"type"`
	UnitAmount     float64   `json:"unit_amount"`
	Quantity       int       `json:"quantity"`
	SubtotalAmount float64   `json:"subtotal_amount"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
