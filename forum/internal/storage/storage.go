package storage

import (
	"database/sql"
)

type Storage struct {
	Auth
	Post
	User
	Comment
	Reaction
	Chat
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Auth: newAuthStorage(db),
		Post: newPostStorage(db),
		User:     newUserStorage(db),
		Comment: newCommentStorage(db),
		Reaction: newReactionStorage(db),
		Chat: newChatStorage(db),
	}
}
