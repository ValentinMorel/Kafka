package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	Consumer_utils "kafka/consumer_utils"
)

func main() {

	// Kafka URL is a []string type
	var KafkaURL []string

	// Parse the argument with flag URL for Kafka broker URL
	KafkaURLArg := flag.String("URL", "localhost:9092", "string")
	// Parse the argument with flag for topic specification
	KafkaTopic := flag.String("topic", "test", "string")
	// Parse the argument with flag for partition specification
	KafkaPartition := flag.Int("partition", 0, "int")
	flag.Parse()

	KafkaURL = append(KafkaURL, *KafkaURLArg)

	broker := Consumer_utils.NewKafkaConsumer(KafkaURL, *KafkaTopic, *KafkaPartition)
	consumer := broker.CreatePartitionConsumer()
	defer consumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	msgCount := 0
	doneCh := make(chan struct{})

	// Thread with goroutine
	go Consumer_utils.MessageListener(consumer, doneCh, signals, msgCount)

	// Sync Channels
	<-doneCh

	fmt.Println("Processed", msgCount, "messages")
}
