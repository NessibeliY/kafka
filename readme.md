# Kafka Producer & Consumer in Go

This repository contains a simple Kafka producer and consumer implemented in Go using the [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go) library.

## Prerequisites

- Kafka installed and running on `localhost:9092`
- Zookeeper running (if using older Kafka versions)
- Go installed (`>=1.18`)
- The `confluent-kafka-go` package installed

## Installation

1. Clone the repository:
   ```sh
   git clone git@github.com:NessibeliY/kafka.git
   cd kafka_producer_consumer
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

## Running Kafka

Ensure that Kafka is running locally:

```sh
brew services start zookeeper
brew services start kafka
```

To create a topic manually (if not already created):
```sh
/opt/homebrew/opt/kafka/bin/kafka-topics.sh --create --topic test-topic --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
```

## Running the Producer

To start the producer and send a message:
```sh
 go run producer.go
```

The producer sends a JSON message to the `test-topic` topic.

## Running the Consumer

To start the consumer and listen for messages:
```sh
 go run consumer.go
```

The consumer subscribes to `test-topic` and prints received messages to the console.



