package application

import (
	"PubNotification/src/notification/application/repositories"
	"PubNotification/src/notification/domain"
	"PubNotification/src/notification/domain/entities"
	"log"
)

type CreateNotification struct {
	notificationRepo    domain.INotification
	serviceNotification repositories.INotificationService
}

func NewCreateNotification(notificationRepo domain.INotification, serviceNotification repositories.INotificationService) *CreateNotification {
	return &CreateNotification{
		notificationRepo:    notificationRepo,
		serviceNotification: serviceNotification,
	}
}

func (c *CreateNotification) Execute(asignature entities.Notification) (entities.Notification, error) {
	created, err := c.notificationRepo.Send(asignature.Asignature)
	if err != nil {
		return entities.Notification{}, err
	}

	err = c.serviceNotification.PublishedEvent("AsignatureCreated", created)
	if err != nil {
		log.Printf("Error notifying about created asignature: %v", err)
		return entities.Notification{}, err
	}

	return created, nil
}
