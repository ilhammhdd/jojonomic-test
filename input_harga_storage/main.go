package main

import (
	"net/http"

	"github.com/Shopify/sarama"
	"github.com/gorilla/mux"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

func init() {
	utils.SetupUtils("input_harga_storage")
	utils.OpenDBConnection()
}

func main() {
	msgChan := make(chan *sarama.ConsumerMessage)
	go func() {
		for msg := range msgChan {
			HandleInputHarga(msg)
		}
	}()
	go utils.Consume("input-harga", 0, sarama.OffsetNewest, msgChan)

	r := mux.NewRouter()
	r.HandleFunc("/last-inserted/harga/id", HandleLastInsertedHarga).Methods(http.MethodGet)
	utils.StartServerMux(r)
}
