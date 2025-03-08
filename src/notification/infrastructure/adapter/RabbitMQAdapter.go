package adapter

import (
	"encoding/json"
	"log"

	"PubNotification/src/notification/domain/entities"
	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQAdapter es el adaptador para conectarse y publicar eventos en RabbitMQ.
type RabbitMQAdapter struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

// NewRabbitMQAdapter crea y retorna un nuevo adaptador de RabbitMQ.
func NewRabbitMQAdapter() (*RabbitMQAdapter, error) {
	conn, err := amqp.Dial("amqp://toledo:12345@35.170.134.124:5672/")
	if err != nil {
		log.Printf("Error conectando a RabbitMQ: %v", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Error abriendo canal: %v", err)
		return nil, err
	}

	_, err = ch.QueueDeclare(
		"asignatures", // nombre de la cola
		true,          // durable
		false,         // auto-eliminar
		false,         // no exclusiva
		false,         // sin esperar
		nil,           // argumentos adicionales
	)
	if err != nil {
		log.Printf("Error declarando la cola: %v", err)
		return nil, err
	}

	return &RabbitMQAdapter{conn: conn, ch: ch}, nil
}

// Conn devuelve la conexión subyacente de RabbitMQ.
func (r *RabbitMQAdapter) Conn() *amqp.Connection {
	return r.conn
}

// Close cierra el canal y la conexión de RabbitMQ.
func (r *RabbitMQAdapter) Close() error {
	if r.ch != nil {
		if err := r.ch.Close(); err != nil {
			log.Printf("Error al cerrar el canal de RabbitMQ: %v", err)
			return err
		}
		log.Println("Canal de RabbitMQ cerrado.")
	}

	if r.conn != nil {
		if err := r.conn.Close(); err != nil {
			log.Printf("Error al cerrar la conexión de RabbitMQ: %v", err)
			return err
		}
		log.Println("Conexión de RabbitMQ cerrada.")
	}

	return nil
}

// PublishEvent publica un evento en la cola de RabbitMQ.
func (r *RabbitMQAdapter) PublishEvent(eventType string, data entities.Notification) error {
	if r.ch == nil || r.ch.IsClosed() {
		var err error
		r.ch, err = r.conn.Channel()
		if err != nil {
			log.Printf("Error reabriendo el canal de RabbitMQ: %v", err)
			return err
		}
	}

	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error convirtiendo el evento a JSON: %v", err)
		return err
	}

	err = r.ch.Publish(
		"",
		"asignatures", // nombre de la cola
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Printf("Error enviando mensaje a RabbitMQ: %v", err)
		return err
	}

	log.Println("Evento publicado:", eventType)
	return nil
}
