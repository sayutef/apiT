package repositories

import (
	"PubNotification/src/notification/domain/entities"
	"log"
)

type PublisherInterface interface {
	PublishEvent(eventType string, data entities.Notification) error
}

type ServiceNotification struct {
	publisher PublisherInterface
}

func (s *ServiceNotification) PublishedEvent(eventType string, notification entities.Notification) error {
	log.Printf("Publishing event %s for notification: %v", eventType, notification)
	return s.publisher.PublishEvent(eventType, notification)
}

func NewServiceNotification(publisher PublisherInterface) INotificationService {
	return &ServiceNotification{publisher: publisher}
}
