package entity

import (
	"time"
)

// Entry represents a financial entry record.
type Entry struct {
	EntryID       int       `json:"entry_id"`
	Sign          string    `json:"sign"`
	Acc           string    `json:"acc"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"`
	TransactionID int       `json:"transaction_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
