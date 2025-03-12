package domain

import "PubNotification/src/notification/domain/entities"

type INotification interface {
	Send(message entities.Notification) error
}
