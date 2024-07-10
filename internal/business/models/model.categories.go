package models

import "time"

type Category struct {
	ID        int       `json:"int"`
	Name      string    `json:"string"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
