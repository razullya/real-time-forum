package storage

import (
	"database/sql"
	"fmt"
	"forum/models"
)

type Auth interface {
	CreateUser(user models.User) error
	GetUserByLogin(login string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	SaveSessionToken(login, token string) error
	GetUserByToken(token string) (models.User, error)
	DeleteSessionToken(token string) error
	CheckToken(token string) error
}

type AuthStorage struct {
	db *sql.DB
}

func newAuthStorage(db *sql.DB) *AuthStorage {
	return &AuthStorage{
		db: db,
	}
}

func (s *AuthStorage) CreateUser(user models.User) error {
	query := `INSERT INTO user(email, username, hashPassword) VALUES ($1, $2, $3);`
	_, err := s.db.Exec(query, user.Email, user.Username, user.Password)
	if err != nil {
		return fmt.Errorf("storage: create user: %w", err)
	}
	return nil
}

func (s *AuthStorage) GetUserByLogin(username string) (models.User, error) {
	query := `SELECT id, email, username, hashPassword FROM user WHERE username=$1;`
	row := s.db.QueryRow(query, username)
	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("storage: get user by login: %w", err)
	}
	return user, nil
}
func (s *AuthStorage) GetUserByEmail(email string) (models.User, error) {
	query := `SELECT id, email, email, hashPassword FROM user WHERE email=$1;`
	row := s.db.QueryRow(query, email)
	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("storage: get user by login: %w", err)
	}
	return user, nil
}
func (s *AuthStorage) SaveSessionToken(username, token string) error {
	query := `UPDATE user SET session_token = $1 WHERE username = $3;`
	_, err := s.db.Exec(query, token, username)
	if err != nil {
		return fmt.Errorf("storage: save session token: %w", err)
	}
	return nil
}

func (s *AuthStorage) GetUserByToken(token string) (models.User, error) {
	query := `SELECT id, email, username, hashPassword FROM user WHERE session_token=$1;`
	row := s.db.QueryRow(query, token)
	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("no user")
	}
	return user, nil
}

func (s *AuthStorage) DeleteSessionToken(token string) error {
	query := `UPDATE user SET session_token = NULL WHERE session_token = $1;`
	_, err := s.db.Exec(query, token)
	if err != nil {
		return fmt.Errorf("storage: delete session token: %w", err)
	}
	return nil
}
func (s *AuthStorage) CheckToken(token string) error {
	query := `SELECT username FROM user WHERE session_token = $1;`
	_, err := s.db.Exec(query, token)
	if err != nil {
		return fmt.Errorf("storage: no token: %w", err)
	}
	return nil
}
