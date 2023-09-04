package kafka

import (
	"encoding/json"
	"go-sarama-kafka/producer/model"
	"log"

	"github.com/IBM/sarama"
)

func KafkaBootstrap() sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForLocal

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %v", err)
	}

	return producer
}

func SendMessage(producer sarama.SyncProducer, message model.User)  {
	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	// Create a Kafka message
	kafkaMessage := &sarama.ProducerMessage{
		Topic: "purchases",
		Value: sarama.StringEncoder(jsonData),
	}

	// Send the message to Kafka
	partition, offset, err := producer.SendMessage(kafkaMessage)
	if err != nil {
		log.Fatalf("Error sending message to Kafka: %v", err)
	}

	log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}