package models

type User struct {
	ID       int    `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
	Posts    []Post `json:"posts,omitempty"`
}
