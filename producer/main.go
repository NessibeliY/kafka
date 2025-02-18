package main

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	kafkaBroker = "localhost:9092"
	topic       = "test-topic"
)

type Message struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaBroker})
	if err != nil {
		fmt.Println("Failed to create producer:", err)
		return
	}
	defer p.Close()

	message := Message{
		Key:   "example_key",
		Value: "Hello, Kafka2!",
	}

	serializedMessage, err := serializeMessage(message)
	if err != nil {
		fmt.Println("Failed to serialize message:", err)
		return
	}

	err = produceMessage(p, topic, serializedMessage)
	if err != nil {
		fmt.Println("Failed to produce message:", err)
		return
	}

	fmt.Println("Successfully produced message")
}

// serialize the message struct to JSON
func serializeMessage(message Message) ([]byte, error) {
	serialized, err := json.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("Failed to serialize message: %v", err)
	}
	return serialized, nil
}

func produceMessage(p *kafka.Producer, topic string, message []byte) error {
	kafkaMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}

	// produce the Kafka message
	deliveryChan := make(chan kafka.Event)
	err := p.Produce(kafkaMessage, deliveryChan)
	if err != nil {
		return fmt.Errorf("Failed to produce message: %v", err)
	}

	// wait for delivery report or error
	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return fmt.Errorf("Delivery failed: %v", m.TopicPartition.Error)
	}

	close(deliveryChan)

	return nil
}
