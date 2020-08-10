package main

import (
	"flag"
	"fmt"
	Producer_utils "kafka/producer_utils"

	"github.com/Shopify/sarama"
)

const (
	KafkaURL   = "localhost:9092"
	KafkaTopic = "test"
)

func main() {

	// Kafka URL is a []string type
	var KafkaURL []string

	// Parse the argument with flag URL for Kafka broker URL
	KafkaURLArg := flag.String("URL", "localhost:9092", "string")
	KafkaURL = append(KafkaURL, *KafkaURLArg)

	// Parse the argument with flag for topic specification
	KafkaTopic := flag.String("topic", "test", "string")

	broker := Producer_utils.NewKafkaProducer(KafkaURL)
	producer, err := broker.CreateProducer()
	defer producer.Close()

	if err != nil {
		//panic is a way to handle unexpected behavior
		panic(err)
	}

	msg := &sarama.ProducerMessage{
		Topic: *KafkaTopic,
		Value: sarama.StringEncoder("Something"),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Message is stored in topic (%s)/partition(%d)/offset(%d)\n", *KafkaTopic, partition, offset)

}
