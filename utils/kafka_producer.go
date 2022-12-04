package utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/Shopify/sarama"
)

type ValueByteConverterFunc[T any] func(t *T) ([]byte, error)

func (vbcf ValueByteConverterFunc[T]) ConvertToByteValue(t *T) ([]byte, error) {
	return vbcf(t)
}

type ValueByteConverter[T any] interface {
	ConvertToByteValue(t *T) ([]byte, error)
}

func Produce[T any](topic string, val *T, valConv ValueByteConverter[T], keyPrefix ...string) (string, error) {
	producer, err := sarama.NewSyncProducer([]string{ENV[ENV_KAFKA_BROKER]}, KafkaClientConfig)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Println(err)
		}
	}()

	key, err := SID.Generate()
	if err != nil {
		return "", err
	}
	key = fmt.Sprintf("%s%s", strings.Join(keyPrefix, ""), key)
	valByte, err := valConv.ConvertToByteValue(val)
	if err != nil {
		return "", err
	}

	msg := &sarama.ProducerMessage{Topic: topic, Key: sarama.StringEncoder(key), Value: sarama.ByteEncoder(valByte)}
	_, _, err = producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
		return "", err
	}
	return key, nil
}

type Production[T any] struct {
	Topic        string
	Val          *T
	ValConv      ValueByteConverter[T]
	syncProducer sarama.SyncProducer
}

func NewProduction[T any](topic string, val *T, valConv ValueByteConverter[T]) (Production[T], error) {
	producer, err := sarama.NewSyncProducer([]string{ENV[ENV_KAFKA_BROKER]}, KafkaClientConfig)
	if err != nil {
		log.Println(err)
		return Production[T]{}, err
	}

	return Production[T]{Topic: topic, Val: val, ValConv: valConv, syncProducer: producer}, nil
}

func (p Production[T]) Prepare() (string, []byte, error) {
	key, err := SID.Generate()
	if err != nil {
		return "", nil, err
	}
	valBytes, err := p.ValConv.ConvertToByteValue(p.Val)
	if err != nil {
		return "", nil, err
	}

	return key, valBytes, nil
}

func (p Production[T]) Produce(key string, valBytes []byte) error {
	msg := &sarama.ProducerMessage{Topic: p.Topic, Key: sarama.StringEncoder(key), Value: sarama.ByteEncoder(valBytes)}
	_, _, err := p.syncProducer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
		return err
	}
	return nil
}
