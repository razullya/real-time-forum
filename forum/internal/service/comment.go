package service

import (
	"forum/internal/storage"
	"forum/models"
	"strconv"
)

type Comment interface {
	CreateComment(id string, token string, comment string) error
	GetCommentsByIdPost(id string) ([]models.Comment, error)
}
type CommentService struct {
	storage storage.Storage
}

func newCommentService(storage *storage.Storage) *CommentService {
	return &CommentService{
		storage: *storage,
	}
}
func (c *CommentService) CreateComment(id string, token string, comment string) error {
	idd, _ := strconv.Atoi(id)
	user, err := c.storage.Auth.GetUserByToken(token)
	if err != nil {
		return err
	}
	if err := c.storage.CreateComment(idd, user.Username, comment); err != nil {
		return err
	}
	return nil
}
func (c *CommentService) GetCommentsByIdPost(id string) ([]models.Comment, error) {
	idd, _ := strconv.Atoi(id)

	comments, err := c.storage.GetCommentsByIdPost(idd)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
