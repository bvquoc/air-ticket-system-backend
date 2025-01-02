package entity

import (
	"time"
)

// Passenger represents a passenger record.
type Passenger struct {
	PassengerID int       `json:"passenger_id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
