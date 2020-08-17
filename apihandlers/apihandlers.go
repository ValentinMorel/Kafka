package apihandlers

import (
	"fmt"
	"kafka/broker"
	"net/http"
	"strconv"

	"github.com/Shopify/sarama"
	"github.com/labstack/echo"
)

// Struct of the messages we can send on the API when calling :
// curl \
//  -i -X POST \
// 	-H 'Content-Type: application/json' \
// 	-d '{"key": "User", "value": "user@email.com"}' \
//   http://localhost:8888/producer/sync

type messagetoqueue struct {
	Key   string `json::"key"`
	Value string `json::"value"`
}

const (
	KafkaUrl       = "localhost:9092"
	KafkaSyncTopic = "test"
	KafkaPartition = "1"
)

func SendSyncMsg(echoContext echo.Context) error {

	var message messagetoqueue

	if err := echoContext.Bind(&message); err != nil {
		return echoContext.NoContent(http.StatusBadGateway)
	}

	kafkabroker := broker.NewKafkaBroker([]string{KafkaUrl}, "", "", false)
	producer, err := kafkabroker.CreateProducer()
	defer producer.Close()

	if err != nil {
		return echoContext.NoContent(http.StatusInternalServerError)
	}

	synctopic := KafkaSyncTopic
	partitionToEnquene, err := strconv.Atoi(KafkaPartition)
	if err != nil {
		partitionToEnquene = 0
	}

	kafkaMessage := sarama.ProducerMessage{
		Topic:     synctopic,
		Key:       sarama.StringEncoder(message.Key),
		Value:     sarama.StringEncoder(message.Value),
		Partition: int32(partitionToEnquene),
	}

	partition, offset, err := producer.SendMessage(&kafkaMessage)
	if err != nil {
		return echoContext.String(http.StatusInternalServerError, "Error to send message")
	}

	return echoContext.String(http.StatusOK, fmt.Sprintf("Message stored in topic %s, partition %d, offset %d", synctopic, partition, offset))
}
