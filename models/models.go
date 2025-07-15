package models

import "time"

type Payment struct {
	ID         int       `json:"id"`
	AccountID  int       `json:"account_id"`
	CategoryID int       `json:"category_id"`
	Amount     float64   `json:"amount"`
	Direction  string    `json:"direction"` // 'I' — приход, 'O' — расход
	PayDate    time.Time `json:"pay_date"`
}
