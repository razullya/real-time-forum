package models

import "time"

type Comment struct {
	Id         int       `json:"id,omitempty"`
	PostId     int       `json:"post_id,omitempty"`
	Creator    string    `json:"creator,omitempty"`
	Text       string    `json:"text,omitempty"`
	Likes      int       `json:"likes,omitempty"`
	Dislikes   int       `json:"dislikes,omitempty"`
	Created_at time.Time `json:"created___at,omitempty"`
}
