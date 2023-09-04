package main

import (
	"go-sarama-kafka/consumer/kafka"
	"log"
)

func main() {
	config := kafka.CreateKafkaConfig()
	consumer, err := kafka.CreateKafkaConsumer(config)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := kafka.CreatePartitionConsumer(consumer, "purchases", 0)
	if err != nil {
		log.Fatalf("Error consuming partition: %v", err)
	}
	defer partitionConsumer.Close()

	kafka.ProcessMessages(partitionConsumer)
}

