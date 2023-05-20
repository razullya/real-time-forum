package service

import (
	"crypto/rand"
	"encoding/base64"
)

type Chat interface {
	GenerateToken() (string, error)
}

type ChatService struct {
}

func newChatService() *ChatService {
	return &ChatService{}
}

func (c *ChatService) GenerateToken() (string, error) {

	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil

}
