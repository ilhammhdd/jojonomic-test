package main

import (
	"encoding/json"

	"github.com/ilhammhdd/jojonomic_test/model"
)

type outCekHarga struct {
	Topup   int64 `json:"harga_topup"`
	Buyback int64 `json:"harga_buyback"`
}

func BuildResponseBody(harga *model.Harga) ([]byte, error) {
	out := outCekHarga{harga.Topup, harga.Buyback}
	resp := model.ResponseTmpl{IsError: false, Data: out}
	respJSON, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	return respJSON, nil
}
