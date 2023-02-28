package models

import "time"

type User struct {
	ID        int       `json:"id,omitempty"`
	Email     string    `json:"email,omitempty"`
	Username  string    `json:"username,omitempty"`
	Password  string    `json:"password,omitempty"`
	Token     string    `json:"token,omitempty"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	Posts     []Post    `json:"posts,omitempty"`
}
