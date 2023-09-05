package kafka

import (
	"encoding/json"
	"fmt"
	"go-sarama-kafka/config"
	"go-sarama-kafka/internal/model"
	"go-sarama-kafka/internal/repository"
	ws "go-sarama-kafka/internal/websocket"
	"log"

	"github.com/IBM/sarama"
	"gorm.io/gorm"
)

func StartKafka(env config.EnvVars, db *gorm.DB)  {
	config := CreateKafkaConfig()
	consumer, err := CreateKafkaConsumer(env.KafkaAddress, env.KafkaPort, config)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := CreatePartitionConsumer(consumer, env.KafkaTopic, 0)
	if err != nil {
		log.Fatalf("Error consuming partition: %v", err)
	}
	defer partitionConsumer.Close()

	ProcessMessages(db, partitionConsumer)
}

func CreateKafkaConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	return config
}

func CreateKafkaConsumer(address string, port string, config *sarama.Config) (sarama.Consumer, error) {
	consumer, err := sarama.NewConsumer([]string{fmt.Sprintf("%s:%s", address, port)}, config)
	return consumer, err
}

func CreatePartitionConsumer(consumer sarama.Consumer, topic string, partition int32) (sarama.PartitionConsumer, error) {
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
	return partitionConsumer, err
}

func ProcessMessages(db *gorm.DB, partitionConsumer sarama.PartitionConsumer) {
	for {
		select {
		case message := <-partitionConsumer.Messages():
			HandleMessage(db, message)
		case err := <-partitionConsumer.Errors():
			log.Printf("Error: %v\n", err.Err)
		}
	}
}

func HandleMessage(db *gorm.DB, message *sarama.ConsumerMessage) {
	var user model.User
	err := json.Unmarshal(message.Value, &user)
	if err != nil {
		log.Printf("Error unmarshaling JSON: %v\n", err)
		return
	}
	log.Print(user)

	err = repository.InsertUser(db, &user)
	if err != nil {
		log.Fatalf("Error inserting user into database: %v", err)
	}

	// received message from kafka, send to websocket
	ws.SendWebSocketUpdate(fmt.Sprintf("Sending message to websocket: %v", user))
	// ws.SendWebSocketUpdate(user.Name)
}