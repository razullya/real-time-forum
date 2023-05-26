package service

import "github.com/google/uuid"

type Chat interface {
	GenerateToken() (string, error)
}

type ChatService struct {
}

func newChatService() *ChatService {
	return &ChatService{}
}

func (c *ChatService) GenerateToken() (string, error) {

	return uuid.NewString(), nil

}
