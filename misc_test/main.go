package main

import (
	kafkazookeeper "github.com/ilhammhdd/jojonomic_test/misc_test/kafka_zookeeper"
)

func main() {
	msgNum := 5
	doneChan := make(chan bool)
	go kafkazookeeper.TestConsume(-1, msgNum, doneChan)
	for i := 0; i < msgNum; i++ {
		kafkazookeeper.TestProduce()
	}
	<-doneChan
}
