package storage

import (
	"database/sql"
)

type Chat interface {
	GetChatByToken(users []string) (string, error)
	SaveTokenByUsers(users []string, token string) (string, error)
}

type ChatStorage struct {
	db *sql.DB
}

func newChatStorage(db *sql.DB) *ChatStorage {
	return &ChatStorage{
		db: db,
	}
}
func (c *ChatStorage) GetChatByToken(users []string) (string, error) {
	query := `SELECT token FROM chat WHERE first_user=$1 AND second_user=$2;`
	row := c.db.QueryRow(query, users[0], users[1])
	token := ""
	err := row.Scan(&token)
	if err != nil {
		return "", err
	}

	return token, nil
}
func (c *ChatStorage) SaveTokenByUsers(users []string, token string) (string, error) {
	query := `INSERT INTO chat(first_user, second_user, token) VALUES($1, $2, $3);`
	if _, err := c.db.Exec(query, users[0], users[1], token); err != nil {
		return "", err
	}
	return token, nil
}
