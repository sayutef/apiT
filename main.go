package main

import (
	"PubNotification/src/notification/infrastructure"
	routes "PubNotification/src/notification/infrastructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	dependencies.Init()
	defer dependencies.Close()

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	routes.NewNotificationRepository()

	r.Run(":8083")
}
