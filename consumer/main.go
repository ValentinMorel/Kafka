package main

import (
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
		log.Fatal("Usage : ./main address topic")
	}

	// Kafka URL is a []string type
	var KafkaURL []string

	KafkaURL = append(KafkaURL, os.Args[1])
	KafkaTopic := os.Args[2]

	broker := Consumer_utils.NewKafkaBroker(KafkaURL)
	master, err := broker.CreateMaster()
	defer master.Close()

	if err != nil {
		panic(err)
	}

	// Assign the broker to a topic
	consumer, err := master.ConsumePartition(KafkaTopic, 0, sarama.OffsetOldest)
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
