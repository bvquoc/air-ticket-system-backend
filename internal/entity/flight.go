package entity

import (
	"time"
)

// Flight represents a flight record.
type Flight struct {
	FlightID           int       `json:"flight_id"`
	AircraftID         int       `json:"aircraft_id"`
	DepartureAirport   string    `json:"departure_airport"`
	DestinationAirport string    `json:"destination_airport"`
	DepartureTime      time.Time `json:"departure_time"`
	DestinationTime    time.Time `json:"destination_time"`
	Status             string    `json:"status"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
