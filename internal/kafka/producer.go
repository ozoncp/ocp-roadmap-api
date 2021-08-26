package kafka

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
	"time"
)

//go:generate mockgen -destination=../mocks/producer_mock.go -package=mocks github.com/ozoncp/ocp-roadmap-api/internal/kafka Producer

type Producer interface {
	Send(message Message) error
}

type producer struct {
	prod  sarama.SyncProducer
	topic string
}

func InitProducer() *producer {
	brokers := []string{"kafka:9092"}

	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	sProducer, err := sarama.NewSyncProducer(brokers, config)

	if err != nil {
		log.Panic().Msgf("failed to connect to Kafka brokers: %v", err)
	}

	return &producer{
		prod:  sProducer,
		topic: "roadmap",
	}
}

func (p *producer) Send(message Message) error {
	bytes, err := json.Marshal(message)
	if err != nil {
		log.Err(err).Msg("failed marshaling message to json:")
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic:     p.topic,
		Key:       sarama.StringEncoder(p.topic),
		Value:     sarama.StringEncoder(bytes),
		Partition: -1,
		Timestamp: time.Time{},
	}

	_, _, err = p.prod.SendMessage(msg)
	return err
}
