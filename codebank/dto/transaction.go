package dto

import "time"

type Transaction struct {
	ID              string
	Name            string
	Number          string
	Country         string
	Currency        string
	ExpirationMonth int32
	ExpirationYear  int32
	CVV             int32
	Amount          float64
	Store           string
	Description     string
	CreatedAt       time.Time
}
