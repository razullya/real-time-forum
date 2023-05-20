package service

import (
	"forum/internal/storage"
	"forum/models"
)

type Notification interface {
	GetNotificationByUsername(username string) ([]models.Notification, error)
	CreateNotification(notify models.Notification) error
}

type NotificationService struct {
	stor *storage.Storage
}

func newNotificationService(stor *storage.Storage) *NotificationService {
	return &NotificationService{stor: stor}
}

func (n *NotificationService) GetNotificationByUsername(username string) ([]models.Notification, error) {
	return n.stor.Notification.GetNotificationByUsername(username)
}

func (n *NotificationService) CreateNotification(notify models.Notification) error {
	return n.stor.Notification.CreateNotification(notify)
}
