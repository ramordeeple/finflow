package models

import "time"

type Cashflow struct {
	AccountName string  `json:"account_name"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
}

type Entry struct {
	ID          int64
	AccountID   int64
	AmountMinor int64
	Currency    string
	IsDebit     bool
	Timestamp   time.Time
}