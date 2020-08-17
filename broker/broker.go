package broker

import (
	"strconv"

	"github.com/Shopify/sarama"
)

type KafkaBroker struct {
	Address        []string
	Config         *sarama.Config
	KafkaTopic     string
	KafkaPartition string
}

func NewKafkaBroker(address []string, kafkatopic string, kafkapartition string, consumer bool) *KafkaBroker {
	var Broker *KafkaBroker
	switch consumer {
	case true:
		Broker = &KafkaBroker{
			Address:        address,
			Config:         createconsumerconfig(),
			KafkaTopic:     kafkatopic,
			KafkaPartition: kafkapartition,
		}
		return Broker

	case false:
		Broker = &KafkaBroker{
			Address:        address,
			Config:         createproducerconfig(),
			KafkaTopic:     "",
			KafkaPartition: "",
		}
		return Broker
	}
	return nil
}

func createconsumerconfig() *sarama.Config {
	Config := sarama.NewConfig()
	Config.Consumer.Return.Errors = true
	return Config
}

func createproducerconfig() *sarama.Config {
	Config := sarama.NewConfig()
	Config.Producer.RequiredAcks = sarama.WaitForAll
	Config.Producer.Return.Successes = true
	Config.Producer.Retry.Max = 5
	return Config
}

func (c *KafkaBroker) CreateConsumerMaster() (sarama.Consumer, error) {
	master, err := sarama.NewConsumer(c.Address, c.Config)
	return master, err
}

func (c *KafkaBroker) CreatePartitionConsumer() sarama.PartitionConsumer {
	master, err := c.CreateConsumerMaster()
	if err != nil {
		panic(err)
	}
	partition, _ := strconv.Atoi(c.KafkaPartition)
	consumer, err := master.ConsumePartition(c.KafkaTopic, int32(partition), sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	return consumer
}

func (c *KafkaBroker) CreateProducer() (sarama.SyncProducer, error) {

	producer, err := sarama.NewSyncProducer(c.Address, c.Config)
	return producer, err
}
