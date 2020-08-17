package main

import (
	"flag"
	"fmt"
	"kafka/broker"

	"github.com/Shopify/sarama"
)

func main() {

	// Kafka URL is a []string type
	var KafkaUrl []string

	// Parse the argument with flag URL for Kafka broker URL
	KafkaUrlArg := flag.String("URL", "localhost:9092", "string")
	// Parse the argument with flag for topic specification
	KafkaTopic := flag.String("topic", "test", "string")
	flag.Parse()

	KafkaUrl = append(KafkaUrl, *KafkaUrlArg)

	// Configuration is not explicit here
	// Could be a file parsing to extract configuration
	KafkaBroker := broker.NewKafkaBroker(KafkaUrl, *KafkaTopic, string(0), false)
	producer, err := KafkaBroker.CreateProducer()
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
