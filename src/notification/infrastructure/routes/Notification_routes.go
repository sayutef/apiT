package repository

import (
	"PubNotification/src/notification/domain"
	"PubNotification/src/notification/domain/entities"
)

type NotificationRepository struct{}

func NewNotificationRepository() domain.INotification {
	return &NotificationRepository{}
}

func (nr *NotificationRepository) Send(asignature string) (entities.Notification, error) {
	notification := entities.Notification{
		ID:         1,
		Asignature: asignature,
		Message:    "La asignatura se ha registrado: " + asignature,
	}
	return notification, nil
}
