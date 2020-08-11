package APIhandlers

import (
	"fmt"
	"kafka/producer_utils"
	"net/http"
	"strconv"

	"github.com/Shopify/sarama"
	"github.com/labstack/echo"
)

// Struct of the messages we can send on the API when calling :
// $ curl -i -X POST -H 'Content-Type: application/json' -d '{"key": "User", "value": "user@email.com"}' http://localhost:8888/producer/sync
type MessageToQueue struct {
	Key   string `json::"key"`
	Value string `json::"value"`
}

const (
	KAFKA_URL        = "localhost:9092"
	KAFKA_SYNC_TOPIC = "test"
	KAFKA_PARTITION  = "1"
)

func SendSyncMsg(echoContext echo.Context) error {

	var message MessageToQueue

	if err := echoContext.Bind(&message); err != nil {
		return echoContext.NoContent(http.StatusBadGateway)
	}

	broker := producer_utils.NewKafkaBroker([]string{KAFKA_URL})
	producer, err := broker.CreateProducer()
	defer producer.Close()

	if err != nil {
		return echoContext.NoContent(http.StatusInternalServerError)
	}

	syncTopic := KAFKA_SYNC_TOPIC
	partitionToEnquene, err := strconv.Atoi(KAFKA_PARTITION)
	if err != nil {
		partitionToEnquene = 0
	}

	kafkaMessage := sarama.ProducerMessage{
		Topic:     syncTopic,
		Key:       sarama.StringEncoder(message.Key),
		Value:     sarama.StringEncoder(message.Value),
		Partition: int32(partitionToEnquene),
	}

	partition, offset, err := producer.SendMessage(&kafkaMessage)
	if err != nil {
		return echoContext.String(http.StatusInternalServerError, "Error to send message")
	}

	return echoContext.String(http.StatusOK, fmt.Sprintf("Message stored in topic %s, partition %d, offset %d", syncTopic, partition, offset))
}
