package main

import (
	"github.com/Shopify/sarama"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

func init() {
	utils.SetupUtils("input_harga_storage")
}

func main() {
	msgChan := make(chan *sarama.ConsumerMessage)
	utils.GoReceiveConsumerMsg(msgChan, HandleInputHarga)
	utils.Consume("input-harga", 0, sarama.OffsetNewest, msgChan)
}
