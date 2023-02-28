package models

type Reaction struct {
	Id       int    `json:"id,omitempty"`
	PostId   int    `json:"post_id,omitempty"`
	Reaction string `json:"reaction,omitempty"`
	Username string `json:"username,omitempty"`
	Object   string `json:"object,omitempty"`
}
