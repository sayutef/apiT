// infrastructure/dependencies.go
package infraestructure

import (
    "PubNotification/src/notification/application"
    "PubNotification/src/notification/application/repositories"
    "PubNotification/src/notification/infrastructure/adapter"
    "PubNotification/src/notification/infrastructure/controllers"
    "log"
)

type DependenciesAsignature struct {
    CreateAsignatureController *controllers.CreateAsignatureController
    RabbitMQAdapter            *adapter.RabbitMQAdapter
}

func InitAsignature() *DependenciesAsignature {
    // Crear el cliente de RabbitMQ
    rmqClient, err := adapter.NewRabbitMQAdapter()
    if err != nil {
        log.Fatalf("Error creating RabbitMQ client: %v", err)
    }

    // Crear el servicio de notificación
    messageService := repositories.NewServiceNotification(rmqClient) // Asegurarse de que messageService implementa INotification

    // Crear el caso de uso de creación de asignatura
    createAsignatureUseCase := application.NewCreateAsignature(messageService, rmqClient) // Cambié el orden para cumplir con los parámetros correctos

    // Retornar las dependencias con los controladores configurados
    return &DependenciesAsignature{
        CreateAsignatureController: controllers.NewCreateAsignatureController(createAsignatureUseCase, messageService, rmqClient),
        RabbitMQAdapter:            rmqClient,
    }
}
