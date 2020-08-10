package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	Consumer_utils "kafka/consumer_utils"

	"github.com/Shopify/sarama"
)

func main() {

	// logging Usage of the program if len of arguments are less than 3
	if len(os.Args) < 3 {
		log.Fatal("Usage : ./main --URL address --topic topic")
	}

	// Kafka URL is a []string type
	var KafkaURL []string

	// Parse the argument with flag URL for Kafka broker URL
	KafkaURLArg := flag.String("URL", "localhost:9092", "string")
	KafkaURL = append(KafkaURL, *KafkaURLArg)

	// Parse the argument with flag for topic specification
	KafkaTopic := flag.String("topic", "test", "string")

	broker := Consumer_utils.NewKafkaBroker(KafkaURL)
	master, err := broker.CreateMaster()
	defer master.Close()

	if err != nil {
		panic(err)
	}

	// Assign the broker to a topic
	consumer, err := master.ConsumePartition(*KafkaTopic, 0, sarama.OffsetOldest)
	defer consumer.Close()
	if err != nil {
		panic(err)
	}

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
