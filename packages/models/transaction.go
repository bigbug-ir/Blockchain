package models

import "time"

type Transaction struct {
	ID        string
	Amount    int
	Currency  string
	From      string
	To        string
	Timestamp time.Time
}
