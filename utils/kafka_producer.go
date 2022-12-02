package utils

import (
	"log"

	"github.com/Shopify/sarama"
)

type ValueByteConverterFunc[T any] func(t *T) ([]byte, error)

func (vbcf ValueByteConverterFunc[T]) ConvertToByteValue(t *T) ([]byte, error) {
	return vbcf(t)
}

type ValueByteConverter[T any] interface {
	ConvertToByteValue(t *T) ([]byte, error)
}

func Produce[T any](topic string, val *T, valConv ValueByteConverter[T]) (string, error) {
	producer, err := sarama.NewSyncProducer([]string{ENV["KAFKA_BROKER"]}, nil)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	key, err := SID.Generate()
	if err != nil {
		return "", err
	}
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
