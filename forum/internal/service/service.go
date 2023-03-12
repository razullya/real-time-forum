package service

import (
	"forum/internal/storage"
)

type Service struct {
	Auth
	Post
	User
	Comment
	Reaction
}

func NewService(storages *storage.Storage) *Service {
	return &Service{
		Auth:     newAuthService(storages.Auth),
		Post:     newPostService(storages.Post),
		User:     newUserService(storages),
		Comment:  newCommentService(storages),
		Reaction: newReactionService(storages.Reaction),
	}
}
