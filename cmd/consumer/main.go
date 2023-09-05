package main

import (
	"go-sarama-kafka/config"
	"go-sarama-kafka/internal/db"
	broker "go-sarama-kafka/internal/kafka"
	ws "go-sarama-kafka/internal/websocket"

	"github.com/gin-gonic/gin"
)

func main() {
	env, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	config := &db.ConfigDB{
		Host:     env.DBHost,
		User:     env.DBUser,
		Password: env.DBPass,
		DBName:   env.DBName,
		Port:     env.DBPort,
		SSLMode:  env.DBSSLmode,
	}

	go func() {
		db := db.ConnectDB(config)

		// Start Kafka consumer
		broker.StartKafka(env, db)
	}()

	r := gin.Default()

	go ws.StartWebSocketServer(env.WsPort)

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	// Websocket endpoint
	r.GET("/ws", ws.HandleWebSocket)

	r.Run(":" + env.Port)
}