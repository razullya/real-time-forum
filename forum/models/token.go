package models

import "time"

type Token struct {
	Token     string    `json:"token,omitempty"`
	ExpiresAT time.Time `json:"expires_at,omitempty"`
}
