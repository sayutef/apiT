package repositories

import "PubNotification/src/notification/domain/entities"

type INotificationService interface {
	PublishedEvent(eventType string, notification entities.Notification) error
}
