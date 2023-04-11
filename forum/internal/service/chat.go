package service

import (
	"database/sql"
	"forum/internal/storage"

	"github.com/google/uuid"
)

type Chat interface {
	GetChatByTokenAndUsername(username string, token string) ([]string, error)
	SaveTokenByUsers(users []string) ([]string, error)
	CheckToken(token []string) error
}

type ChatService struct {
	storage *storage.Storage
}

func newChatService(storage *storage.Storage) *ChatService {
	return &ChatService{
		storage: storage,
	}
}
func (c *ChatService) GetChatByTokenAndUsername(username string, token string) ([]string, error) {
	user, err := c.storage.Auth.GetUserByToken(token)
	if err != nil {
		return []string{}, err
	}
	tokens, err := c.storage.Chat.GetTokenByUsers([]string{username, user.Username})
	if err != nil {
		if err.Error() != sql.ErrNoRows.Error() {
			return []string{}, err
		}
		return c.SaveTokenByUsers([]string{username, user.Username})
	}
	return tokens, err
}
func (c *ChatService) SaveTokenByUsers(users []string) ([]string, error) {
	return c.storage.Chat.SaveTokenByUsers(users, []string{uuid.NewString(), uuid.NewString()})
}
func (c *ChatService) CheckToken(token []string) error {
	return c.storage.Chat.CheckToken(token)
}
