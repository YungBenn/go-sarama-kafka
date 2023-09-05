package main

import (
	"go-sarama-kafka/config"
	"go-sarama-kafka/internal/kafka"
	"go-sarama-kafka/internal/model"
)

func main() {
	env, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	producer := kafka.KafkaBootstrap(env.KafkaAddress, env.KafkaPort)
	defer producer.Close()

	message := model.User{
		Name:  "Ruben",
		Age:   22,
		Hobby: "Watch movie",
	}	
	
	kafka.SendMessage(env.KafkaTopic, producer, message)
}
