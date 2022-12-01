package kafkazookeeper

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

func TestConsume(offset int64, breaker int, doneChan chan bool) {
	config := sarama.NewConfig()
	config.ClientID = "misc-test"
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(TOPIC, 0, offset)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message topic: %s, partition: %d, offset: %d, timestamp: %v, headers: %v, key: %s, value: %s\n", msg.Topic, msg.Partition, msg.Offset, msg.Timestamp, msg.Headers, msg.Key, msg.Value)
			consumed++
			if consumed == breaker {
				doneChan <- true
				break ConsumerLoop
			}
		case <-signals:
			break ConsumerLoop
		}
	}
}
