package controllers

import (
	"PubNotification/src/notification/application"
	"PubNotification/src/notification/domain"
	"PubNotification/src/notification/domain/entities"
	"github.com/gin-gonic/gin"
	"PubNotification/src/notification/infrastructure/adapter"
)

type CreateAsignatureController struct {
	useCase    *application.CreateAsignature
	asignature domain.INotification
	wsClient   *adapter.RabbitMQAdapter // Agregar cliente WebSocket
}

func NewCreateAsignatureController(useCase *application.CreateAsignature, asignature domain.INotification, wsClient *adapter.RabbitMQAdapter) *CreateAsignatureController {
	return &CreateAsignatureController{useCase: useCase, asignature: asignature, wsClient: wsClient}
}

func (cs_a *CreateAsignatureController) Execute(c *gin.Context) {
	var asignature entities.Notification
	if err := c.ShouldBindJSON(&asignature); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}

	err := cs_a.useCase.Execute(asignature)
	if err != nil {
		c.JSON(500, gin.H{"error": "No se pudo crear la asignatura"})
		return
	}

	// Emitir el mensaje a través de WebSocket después de que la asignatura haya sido creada
	message := "Asignatura registrada correctamente: " + asignature.Asignature
	cs_a.wsClient.Send(asignature) // Aquí enviaríamos la notificación a través de RabbitMQ (si lo tienes configurado)

	// Notificar en WebSocket
	cs_a.wsClient.PublishEvent("asignatura_creada", asignature)

	c.JSON(200, gin.H{"message": message, "asignature": asignature})
}
