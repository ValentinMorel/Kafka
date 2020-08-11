package consumer_utils

import (
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

func CreateDefaultConfig() *sarama.Config {
	Config := sarama.NewConfig()
	Config.Consumer.Return.Errors = true
	return Config
}

type KafkaBroker struct {
	Address        []string
	Config         *sarama.Config
	KafkaTopic     string
	KafkaPartition int
}

func NewKafkaBroker(address []string, kafkaTopic string, kafkaPartition int) *KafkaBroker {
	return &KafkaBroker{
		Address:        address,
		Config:         CreateDefaultConfig(),
		KafkaTopic:     kafkaTopic,
		KafkaPartition: kafkaPartition,
	}
}

func (c *KafkaBroker) CreateMaster() (sarama.Consumer, error) {
	master, err := sarama.NewConsumer(c.Address, c.Config)
	return master, err
}

func (c *KafkaBroker) CreatePartitionConsumer() sarama.PartitionConsumer {
	master, err := c.CreateMaster()
	if err != nil {
		panic(err)
	}
	consumer, err := master.ConsumePartition(c.KafkaTopic, int32(c.KafkaPartition), sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	return consumer
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
