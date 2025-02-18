package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	kafkaBroker = "localhost:9092"
	topic       = "test-topic"
	groupID     = "test-group"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaBroker,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		fmt.Println("Failed to create consumer:", err)
		return
	}
	defer c.Close()

	// subscribe to the Kafka topic
	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		fmt.Println("Failed to subscribe to topic:", err)
		return
	}

	// setup a channel to handle OS signals for graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	// start consuming messages
	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Received signal %v: terminating\n", sig)
			run = false
		default:
			// poll for Kafka messages
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("Received message from topic %s: %s\n", *e.TopicPartition.Topic, string(e.Value))
			case kafka.Error:
				fmt.Printf("Error from topic: %v\n", e)
			}
		}
	}
}
