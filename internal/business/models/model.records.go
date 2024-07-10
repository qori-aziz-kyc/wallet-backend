package models

import "time"

type Record struct {
	ID         int       `json:"int"`
	UserID     int       `json:"user_id"`
	CategoryID int       `json:"category_id"`
	Amount     float64   `json:"amount"`
	Date       string    `json:"date"`
	DateInTime time.Time `json:"-"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
