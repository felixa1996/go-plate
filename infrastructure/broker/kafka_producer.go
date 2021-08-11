package broker

import (
	"context"
	"net/http"
	"os"

	"github.com/felixa1996/go-plate/adapter/api/logging"
	"github.com/felixa1996/go-plate/adapter/logger"
	"github.com/segmentio/kafka-go"
)

const (
	partition = 0
)

type KafkaProducer struct {
	Ctx    context.Context
	Log    logger.Logger
	LogKey string
	Topic  string
	Key    string
	Value  string
}

func Produce(k *KafkaProducer) error {
	broker1Address := os.Getenv("KAFKA_BROKER_1")
	broker2Address := os.Getenv("KAFKA_BROKER_2")
	broker3Address := os.Getenv("KAFKA_BROKER_3")

	w := &kafka.Writer{
		Addr:      kafka.TCP(broker1Address, broker2Address, broker3Address),
		Topic:     k.Topic,
		Balancer:  &kafka.Hash{},
		BatchSize: 1,
	}

	err := w.WriteMessages(k.Ctx,
		// NOTE: Each Message has Topic defined, otherwise an error is returned.
		kafka.Message{
			Partition: 0,
			// Topic: "user_deleted",
			Key:   []byte(k.Key),
			Value: []byte(k.Value),
		},
	)

	if err != nil {
		logging.NewError(k.Log, err, k.LogKey, http.StatusInternalServerError).Log("Kafka producer error send")
		panic(err)
		// return err
	}

	if err := w.Close(); err != nil {
		logging.NewError(k.Log, err, k.LogKey, http.StatusInternalServerError).Log("Kafka Writer failed to close")
		panic(err)
		// return err
	}
	logging.NewInfo(k.Log, k.LogKey, http.StatusOK).Log("Success send data " + k.LogKey)
	return nil
}
