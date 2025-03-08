package dependencies

import (
	"database/sql"
	"log"

	"PubNotification/src/notification/infrastructure/adapter"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB            *sql.DB
	RabbitAdapter *adapter.RabbitMQAdapter
)

func Init() {
	log.Println("Inicializando dependencias...")

	var err error
	if err != nil {
		log.Fatalf("Error al inicializar MySQL: %v", err)
	}
	log.Println("Conexión a MySQL establecida.")

	RabbitAdapter, err = adapter.NewRabbitMQAdapter()
	if err != nil {
		log.Fatalf("Error al inicializar RabbitMQ: %v", err)
	}
	log.Println("Adaptador RabbitMQ inicializado correctamente.")
}

func Close() {
	log.Println("Cerrando dependencias...")

	if RabbitAdapter != nil && RabbitAdapter.Conn() != nil {
		if err := RabbitAdapter.Conn().Close(); err != nil {
			log.Printf("Error al cerrar conexión de RabbitMQ: %v", err)
		} else {
			log.Println("Conexión de RabbitMQ cerrada.")
		}
	}

	if DB != nil {
		if err := DB.Close(); err != nil {
			log.Printf("Error al cerrar la conexión MySQL: %v", err)
		} else {
			log.Println("Conexión MySQL cerrada.")
		}
	}
}
