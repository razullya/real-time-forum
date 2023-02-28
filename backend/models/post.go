package models

import "time"

type Post struct {
	Id          int       `json:"id,omitempty"`
	Creator     string    `json:"creator,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Category    []string  `json:"category,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	Likes       int       `json:"likes,omitempty"`
	Dislikes    int       `json:"dislikes,omitempty"`
}
