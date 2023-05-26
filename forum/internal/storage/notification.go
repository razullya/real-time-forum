package storage

import (
	"database/sql"
	"fmt"
	"forum/models"
)

type Notification interface {
	GetNotificationByUsername(username string) ([]models.Notification, error)
	CreateNotification(notify models.Notification) error
	NotificationChecked(username string, sender string) error
	GetTokenChat(username string, sender string) (models.Notification, error)
}

type NotificationStorage struct {
	db *sql.DB
}

func newNotificationStorage(db *sql.DB) *NotificationStorage {
	return &NotificationStorage{db: db}
}

func (n *NotificationStorage) GetNotificationByUsername(username string) ([]models.Notification, error) {
	query := `SELECT username, sender, msg, checked FROM notification WHERE username=$1;`
	row, err := n.db.Query(query, username)
	if err != nil {
		return nil, err
	}
	var allNotif []models.Notification
	for row.Next() {
		var notif models.Notification
		if err := row.Scan(&notif.Username, &notif.Sender, &notif.Message, &notif.Checked); err != nil {
			return nil, err
		}
		allNotif = append(allNotif, notif)
	}

	return allNotif, nil
}
func (n *NotificationStorage) CreateNotification(notify models.Notification) error {
	query := `INSERT INTO notification(username, sender, msg,checked) VALUES($1, $2, $3, $4);`
	if _, err := n.db.Exec(query, notify.Username, notify.Sender, notify.Message, false); err != nil {
		return err
	}
	return nil
}
func (n *NotificationStorage) NotificationChecked(username string, sender string) error {
	query := `UPDATE notification SET checked = $1 WHERE username=$2 and sender=$3;`
	_, err := n.db.Exec(query, true, username, sender)
	if err != nil {
		return fmt.Errorf("storage: delete post: %w", err)
	}
	return nil
}
func (n *NotificationStorage) GetTokenChat(username string, sender string) (models.Notification, error) {
	query := `SELECT username, sender, msg, checked FROM notification WHERE username=$1 and sender=$2;`
	fmt.Println(username, sender)
	rows := n.db.QueryRow(query, username, sender)

	var notif models.Notification
	if err := rows.Scan(&notif.Username, &notif.Sender, &notif.Message, &notif.Checked); err != nil {
		return models.Notification{}, err
	}

	return notif, nil
}
