package storage

import (
	"database/sql"
	"forum/models"
)

type Notification interface {
	GetNotificationByUsername(username string) ([]models.Notification, error)
	CreateNotification(notify models.Notification) error
}

type NotificationStorage struct {
	db *sql.DB
}

func newNotificationStorage(db *sql.DB) *NotificationStorage {
	return &NotificationStorage{db: db}
}

func (n *NotificationStorage) GetNotificationByUsername(username string) ([]models.Notification, error) {
	query := `SELECT username, sender, type_msg, msg,checked WHERE username=$1;`
	row, err := n.db.Query(query, username)
	if err != nil {
		return nil, err
	}
	var allNotif []models.Notification
	for row.Next() {
		var notif models.Notification
		if err := row.Scan(&notif.Username, &notif.Sender, &notif.Type, &notif.Message, &notif.Checked); err != nil {
			return nil, err
		}
		allNotif = append(allNotif, notif)
	}

	return allNotif, nil
}
func (n *NotificationStorage) CreateNotification(notify models.Notification) error {
	query := `INSERT INTO notification(username, sender, type_msg, msg) VALUES($1, $2, $3, $4);`
	if _, err := n.db.Exec(query, notify.Username, notify.Sender, notify.Type, notify.Message); err != nil {
		return err
	}
	return nil
}
