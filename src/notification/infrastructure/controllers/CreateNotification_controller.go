package controllers

import (
	"PubNotification/src/notification/application"
	"PubNotification/src/notification/domain/entities"
	"encoding/json"
	"log"
	"net/http"
)

// NotificationController gestiona las solicitudes HTTP para las notificaciones.
type NotificationController struct {
	service *application.CreateNotification
}

// NewNotificationController crea una instancia de NotificationController.
func NewNotificationController(service *application.CreateNotification) *NotificationController {
	return &NotificationController{service: service}
}

// SendNotification procesa la solicitud para enviar una notificación.
func (nc *NotificationController) SendNotification(w http.ResponseWriter, r *http.Request) {
	var body map[string]string
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	message, ok := body["message"]
	if !ok {
		http.Error(w, "Message not provided", http.StatusBadRequest)
		return
	}

	notification := entities.Notification{
		Asignature: message,
	}

	_, err = nc.service.Execute(notification)
	if err != nil {
		log.Printf("Error sending notification: %v", err)
		http.Error(w, "Error sending notification", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Notificación enviada"))
}
