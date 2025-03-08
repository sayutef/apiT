package domain

import "PubNotification/src/notification/domain/entities"

type INotification interface {
	Send(asignature string) (entities.Notification, error)
}
