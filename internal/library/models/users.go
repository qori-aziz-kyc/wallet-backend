package models

import "time"

type User struct {
	ID        int       `json:"int"`
	UserName  string    `json:"username"`
	Password  string    `json:"password"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
