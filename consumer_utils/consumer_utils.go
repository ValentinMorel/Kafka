package consumer_utils

import (
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

func MessageListener(consumer sarama.PartitionConsumer, donech chan struct{}, signals chan os.Signal, msgcount int) {
	for {
		select {
		case err := <-consumer.Errors():
			fmt.Println(err)
		case msg := <-consumer.Messages():
			msgcount++
			fmt.Println("Received messages", string(msg.Key), string(msg.Value))
		case <-signals:
			fmt.Println("Interrupt is detected")
			donech <- struct{}{}
		}
	}
}
