// repositories/ServiceNotification.go
package repositories

import (
    "PubNotification/src/notification/domain/entities"
    "log"
)

type ServiceNotification struct {
    imageService IMessageService
}

// Asegúrate de que ServiceNotification implemente la interfaz INotification
func (sn *ServiceNotification) Send(message entities.Notification) error {
    log.Println("Sending notification:", message.Message)

    // Usa el servicio imageService para publicar el evento (este es tu método PublishEvent)
    err := sn.imageService.PublishEvent("AsignatureCreated", message)
    if err != nil {
        log.Printf("Error sending notification: %v", err)
        return err
    }

    return nil
}

// Crear una nueva instancia de ServiceNotification
func NewServiceNotification(imageService IMessageService) *ServiceNotification {
    return &ServiceNotification{imageService: imageService}
}
