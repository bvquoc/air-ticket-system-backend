package entity

import (
	"time"
)

// BDTicket represents a mapping between booking direction and ticket.
type BDTicket struct {
	BdID      int       `json:"bd_id"`
	TicketID  int       `json:"ticket_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
