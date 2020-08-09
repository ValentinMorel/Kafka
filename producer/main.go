package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

const (
	KafkaURL   = "localhost:9092"
	KafkaTopic = "test"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5

	brokers := []string{KafkaURL}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		//panic is a way to handle unexpected behavior
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	topic := KafkaTopic
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("Something to test broker"),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Message is stored in topic (%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

}
