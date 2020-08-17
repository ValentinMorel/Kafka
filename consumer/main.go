package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"kafka/broker"
	"kafka/consumer_utils"
)

func main() {

	// Kafka URL is a []string type.
	var KafkaUrl []string

	// Parse the argument with flag URL for Kafka broker URL.
	KafkaUrlArg := flag.String("URL", "localhost:9092", "string")

	// Parse the argument with flag for topic specification.
	KafkaTopic := flag.String("topic", "test", "string")

	// Parse the argument with flag for partition specification.
	// Need to Create a partition if the partition != 0.
	KafkaPartition := flag.String("partition", "0", "string")

	flag.Parse()

	KafkaUrl = append(KafkaUrl, *KafkaUrlArg)

	KafkaBroker := broker.NewKafkaBroker(KafkaUrl, *KafkaTopic, *KafkaPartition, true)
	consumer := KafkaBroker.CreatePartitionConsumer()
	defer consumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	msgcount := 0
	donech := make(chan struct{})

	// Thread with goroutine.
	go consumer_utils.MessageListener(consumer, donech, signals, msgcount)

	// Sync Channels.
	<-donech

	fmt.Println("Processed", msgcount, "messages")
}
