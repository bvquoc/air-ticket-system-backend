package entity

import (
	"time"
)

// Booking represents a booking record.
type Booking struct {
	BookingID      int       `json:"booking_id"`
	Status         string    `json:"status"`
	TicketAmount   float64   `json:"ticket_amount"`
	FeeAmount      float64   `json:"fee_amount"`
	DiscountAmount float64   `json:"discount_amount"`
	TotalAmount    float64   `json:"total_amount"`
	CheckoutAt     time.Time `json:"checkout_at"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
