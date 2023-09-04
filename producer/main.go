package main

import (
	"go-sarama-kafka/producer/kafka"
	"go-sarama-kafka/producer/model"
)

func main() {
	producer := kafka.KafkaBootstrap()
	defer producer.Close()

	message := model.User{
		Name:  "Ahmad",
		Age:   12,
		Hobby: "Coding",
	}	
	
	kafka.SendMessage(producer, message)
}
