// repositories/ServiceNotification.go
package repositories

import (
    "PubNotification/src/notification/domain/entities"
    "log"
)

type ServiceNotification struct {
    imageService IMessageService
}


func (sn *ServiceNotification) Send(message entities.Notification) error {
    log.Println("Sending notification:", message.Message)


    err := sn.imageService.PublishEvent("AsignatureCreated", message)
    if err != nil {
        log.Printf("Error sending notification: %v", err)
        return err
    }

    return nil
}

func NewServiceNotification(imageService IMessageService) *ServiceNotification {
    return &ServiceNotification{imageService: imageService}
}
