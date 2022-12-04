package main

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/ilhammhdd/jojonomic_test/model"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

func init() {
	utils.SetupUtils("topup-storage")
	utils.OpenDBConnection()
}

func main() {
	msgChan := make(chan *sarama.ConsumerMessage)
	go func() {
		for msg := range msgChan {
			if !strings.Contains(string(msg.Key), "topup") {
				continue
			}
			var txWithRel model.TransaksiWithRel
			err := json.Unmarshal(msg.Value, &txWithRel)
			if err != nil {
				log.Println(err)
				continue
			}

			err = HandleTransaksi(&txWithRel, TopupDBOperator{Ctx: context.Background()})
			if err != nil {
				log.Println(err)
			}
		}
	}()
	utils.Consume("topup", 0, sarama.OffsetNewest, msgChan)
}
