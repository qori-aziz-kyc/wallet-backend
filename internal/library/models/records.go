package models

import "time"

type Record struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	CategoryID int       `json:"category_id"`
	Amount     float64   `json:"amount"`
	Date       string    `json:"date"`
	DateInTime time.Time `json:"-" gorm:"-"`
	Type       string    `json:"type"`
	Note       string    `json:"note"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
