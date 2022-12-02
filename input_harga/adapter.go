package main

import (
	"encoding/json"

	"github.com/ilhammhdd/jojonomic_test/model"
)

func ValueByteConvertion(t *Request) ([]byte, error) {
	modelVal := model.Harga{AdminID: t.AdminID, Topup: t.HargaTopup, Buyback: t.HargaBuyback}
	val, err := json.Marshal(modelVal)
	if err != nil {
		return nil, err
	}
	return val, nil
}
