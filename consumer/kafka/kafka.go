package kafka

import (
	"encoding/json"
	"fmt"
	"go-sarama-kafka/config"
	"go-sarama-kafka/consumer/model"
	ws "go-sarama-kafka/consumer/websocket"
	"log"

	"github.com/IBM/sarama"
)

func StartKafka(env config.EnvVars)  {
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

	ProcessMessages(partitionConsumer)
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

	// received message from kafka, send to websocket
	ws.SendWebSocketUpdate(user.Name)
}