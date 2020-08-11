package producer_utils

import (
	"github.com/Shopify/sarama"
)

type KafkaBroker struct {
	Address []string
	Config  *sarama.Config
}

func NewKafkaBroker(address []string) *KafkaBroker {
	return &KafkaBroker{
		Address: address,
		Config:  CreateDefaultConfig(),
	}
}

func CreateDefaultConfig() *sarama.Config {
	Config := sarama.NewConfig()
	Config.Producer.RequiredAcks = sarama.WaitForAll
	Config.Producer.Return.Successes = true
	Config.Producer.Retry.Max = 5
	return Config
}

func (c *KafkaBroker) CreateProducer() (sarama.SyncProducer, error) {

	producer, err := sarama.NewSyncProducer(c.Address, c.Config)
	return producer, err
}
