package Consumer_utils

import (
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

type KafkaBroker struct {
	Address []string
	Config  *sarama.Config
}

func NewKafkaConsumer(address []string) *KafkaBroker {
	return &KafkaBroker{
		Address: address,
		Config:  CreateDefaultConfig(),
	}
}

func CreateDefaultConfig() *sarama.Config {
	Config := sarama.NewConfig()
	Config.Consumer.Return.Errors = true
	return Config
}

func (c *KafkaBroker) CreateMaster() (sarama.Consumer, error) {

	master, err := sarama.NewConsumer(c.Address, c.Config)
	return master, err
}

func MessageListener(consumer sarama.PartitionConsumer, doneCh chan struct{}, signals chan os.Signal, msgCount int) {
	for {
		select {
		case err := <-consumer.Errors():
			fmt.Println(err)
		case msg := <-consumer.Messages():
			msgCount++
			fmt.Println("Received messages", string(msg.Key), string(msg.Value))
		case <-signals:
			fmt.Println("Interrupt is detected")
			doneCh <- struct{}{}
		}
	}
}
