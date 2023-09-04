package kafka

import (
	"encoding/json"
	"go-sarama-kafka/consumer/model"
	"log"

	"github.com/IBM/sarama"
)

func CreateKafkaConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	return config
}

func CreateKafkaConsumer(config *sarama.Config) (sarama.Consumer, error) {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	return consumer, err
}

func CreatePartitionConsumer(consumer sarama.Consumer, topic string, partition int32) (sarama.PartitionConsumer, error) {
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
	return partitionConsumer, err
}

func ProcessMessages(partitionConsumer sarama.PartitionConsumer) {
	for {
		select {
		case message := <-partitionConsumer.Messages():
			HandleMessage(message)
		case err := <-partitionConsumer.Errors():
			log.Printf("Error: %v\n", err.Err)
		}
	}
}

func HandleMessage(message *sarama.ConsumerMessage) {
	var user model.User
	err := json.Unmarshal(message.Value, &user)
	if err != nil {
		log.Printf("Error unmarshaling JSON: %v\n", err)
		return
	}
	log.Print(user)

	// Lakukan pemrosesan pesan di sini sesuai kebutuhan Anda
}