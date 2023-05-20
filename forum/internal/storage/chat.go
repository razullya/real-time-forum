package storage

import "database/sql"

type Chat interface{
	
}

type ChatStorage struct {
	db *sql.DB
}

func newChatStorage(db *sql.DB) *ChatStorage {
	return &ChatStorage{
		db: db,
	}
}
