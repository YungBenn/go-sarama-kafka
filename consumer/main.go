package main

import (
	"go-sarama-kafka/config"
	broker "go-sarama-kafka/consumer/kafka"
	ws "go-sarama-kafka/consumer/websocket"

	"github.com/gin-gonic/gin"
)

func main() {
	env, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	go ws.StartWebSocketServer(env.WsPort)

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	// Websocket endpoint
	r.GET("/ws", ws.HandleWebSocket)

	// Kafka consumer
	go broker.StartKafka(env)

	r.Run(":" + env.Port)
}
