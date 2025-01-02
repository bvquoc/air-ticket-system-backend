package entity

import (
	"time"
)

// Transaction represents a transaction record.
type Transaction struct {
	TxnID         int       `json:"txn_id"`
	Status        string    `json:"status"`
	Type          string    `json:"type"`
	DebitAcc      string    `json:"debit_acc"`
	CreditAcc     string    `json:"credit_acc"`
	PaymentMethod string    `json:"payment_method"`
	FeeCharge     float64   `json:"fee_charge"`
	Amount        float64   `json:"amount"`
	BookingID     int       `json:"booking_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
