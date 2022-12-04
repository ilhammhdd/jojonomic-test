package main

import (
	"encoding/json"

	"github.com/ilhammhdd/jojonomic_test/model"
)

type requestBody struct {
	Norek string `json:"norek"`
}

type responseBody struct {
	Norek string  `json:"norek"`
	Saldo float64 `json:"saldo"`
}

func BuildResponseBody(rekening *model.Rekening) ([]byte, error) {
	var respBody model.ResponseTmpl
	if rekening == nil {
		respBody = model.ResponseTmpl{IsError: false, Error: ErrNoAccountExists}
	} else {
		respBody = model.ResponseTmpl{IsError: false, Data: responseBody{Norek: rekening.Norek, Saldo: rekening.Saldo}}
	}
	respBodyJSON, err := json.Marshal(respBody)
	if err != nil {
		return nil, err
	}
	return respBodyJSON, err
}
