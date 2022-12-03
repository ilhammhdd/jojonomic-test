package main

import (
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
	"github.com/ilhammhdd/jojonomic_test/model"
)

func HandleInputHarga(consMsg *sarama.ConsumerMessage) {
	var harga model.Harga
	err := json.Unmarshal(consMsg.Value, &harga)
	if err != nil {
		log.Println(err)
		return
	}
	lastInsertedID, err := InsertHarga(&harga)
	if err != nil {
		log.Println(err)
	}
	lastInsertedHargaID.ID = lastInsertedID
}
