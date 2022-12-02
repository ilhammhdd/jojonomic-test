package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/ilhammhdd/jojonomic_test/model"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

func HandleInputHarga(w http.ResponseWriter, r *http.Request) {
	reqBody, err := utils.ReadRequestBodyJSON[Request](r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var errMsgs []string
	if reqBody.AdminID == "" {
		errMsgs = append(errMsgs, "admin_id can't be empty")
	}
	if reqBody.HargaTopup <= 0 {
		errMsgs = append(errMsgs, "harga_topup must be > 0")
	}
	if reqBody.HargaBuyback <= 0 {
		errMsgs = append(errMsgs, "harga_buyback must be > 0")
	}
	if len(errMsgs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		response := model.ResponseTmpl{IsError: true, Error: errors.New(strings.Join(errMsgs, ", "))}
		responseJSON, err := json.Marshal(response)
		if err != nil {
			log.Println(err.Error())
			return
		}
		w.Write(responseJSON)
		return
	}

	key, err := utils.Produce[Request]("input-harga", reqBody, utils.ValueByteConverterFunc[Request](ValueByteConvertion))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	response := model.ResponseTmpl{IsError: false, ReffID: key}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
