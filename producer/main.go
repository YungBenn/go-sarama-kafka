package main

import (
	"go-sarama-kafka/config"
	"go-sarama-kafka/producer/kafka"
	"go-sarama-kafka/producer/model"
)

func main() {
	env, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	producer := kafka.KafkaBootstrap(env.KafkaAddress, env.KafkaPort)
	defer producer.Close()

	message := model.User{
		Name:  "Ahmad",
		Age:   12,
		Hobby: "Coding",
	}	
	
	kafka.SendMessage(env.KafkaTopic, producer, message)
}
