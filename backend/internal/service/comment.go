package service

import (
	"forum/internal/storage"
	"forum/models"
)

type Comment interface {
	CreateComment(id int, username string, comment string) error
	GetCommentsByIdPost(id int) ([]models.Comment, error)
}
type CommentService struct {
	storage storage.Comment
}

func newCommentService(storage storage.Comment) *CommentService {
	return &CommentService{
		storage: storage,
	}
}
func (c *CommentService) CreateComment(id int, username string, comment string) error {
	if err := c.storage.CreateComment(id, username, comment); err != nil {
		return err
	}
	return nil
}
func (c *CommentService) GetCommentsByIdPost(id int) ([]models.Comment, error) {
	comments, err := c.storage.GetCommentsByIdPost(id)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
