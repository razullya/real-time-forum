package storage

import (
	"database/sql"
	"fmt"
)

type Chat interface {
	GetTokenByUsers(users []string) ([]string, error)
	SaveTokenByUsers(users []string, token []string) ([]string, error)
	CheckToken(token []string) error
}

type ChatStorage struct {
	db *sql.DB
}

func newChatStorage(db *sql.DB) *ChatStorage {
	return &ChatStorage{
		db: db,
	}
}
func (c *ChatStorage) GetTokenByUsers(users []string) ([]string, error) {
	query := `SELECT first_token, second_token FROM chat WHERE first_user=$1 AND second_user=$2;`
	row := c.db.QueryRow(query, users[0], users[1])
	firstToken := ""
	secondToken := ""
	err := row.Scan(&firstToken, &secondToken)
	if err != nil {
		return []string{}, err
	}

	return []string{firstToken, secondToken}, nil
}
func (c *ChatStorage) SaveTokenByUsers(users []string, token []string) ([]string, error) {
	query := `INSERT INTO chat(first_user, second_user, first_token, second_token) VALUES($1, $2, $3, $4);`
	if _, err := c.db.Exec(query, users[0], users[1], token[0], token[1]); err != nil {
		return []string{}, fmt.Errorf("no chat")
	}
	return token, nil
}
func (c *ChatStorage) CheckToken(token []string) error {
	query := `SELECT first_user FROM chat WHERE first_token=$1 AND second_person=$2;`
	row := c.db.QueryRow(query, token[0], token[1])

	if len(token) == 1 {
		query = `SELECT first_user FROM chat WHERE second_person=$2;`
		row = c.db.QueryRow(query, token[0])

	}
	var user string
	err := row.Scan(&user)
	if err != nil {
		return err
	}
	if user == "" {
		return fmt.Errorf("no token")
	}
	return nil
}
