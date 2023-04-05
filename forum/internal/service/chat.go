package service

import (
	"database/sql"
	"fmt"
	"forum/internal/storage"

	"github.com/google/uuid"
)

type Chat interface {
	GetChatByToken(username string, token string) (string, error)
	SaveTokenByUsers(users []string) (string, error)
}

type ChatService struct {
	storage *storage.Storage
}

func newChatService(storage *storage.Storage) *ChatService {
	return &ChatService{
		storage: storage,
	}
}
func (c *ChatService) GetChatByToken(username string, token string) (string, error) {
	user, err := c.storage.Auth.GetUserByToken(token)
	if err != nil {
		return "", err
	}
	token, err = c.storage.Chat.GetChatByToken([]string{user.Username, username})
	if err != nil {
		fmt.Println(err.Error(), sql.ErrNoRows.Error())
		if err.Error() != sql.ErrNoRows.Error() {
			return "", err
		}
		return c.SaveTokenByUsers([]string{username, user.Username})
	}
	return token, err
}
func (c *ChatService) SaveTokenByUsers(users []string) (string, error) {
	return c.storage.Chat.SaveTokenByUsers(users, uuid.NewString())
}
