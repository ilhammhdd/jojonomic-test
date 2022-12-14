package utils

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

func Consume(topic string, partition int32, offset int64, msgChan chan *sarama.ConsumerMessage) {
	consumer, err := sarama.NewConsumer([]string{ENV[ENV_KAFKA_BROKER]}, KafkaClientConfig)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Println(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(topic, partition, offset)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Println(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var consumed uint32 = 0
	defer log.Printf("Client %s consumed %d messages from topic: %s, partition: %d, started at offset: %d", KafkaClientConfig.ClientID, consumed, topic, partition, offset)
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			msgChan <- msg
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}
}
