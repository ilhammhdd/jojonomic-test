package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/ilhammhdd/jojonomic_test/model"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

func HandleInputHarga(w http.ResponseWriter, r *http.Request) {
	reqBody, err := utils.ReadRequestBodyJSON[Request](r)
	if utils.IsNotNilAndWriteErrResp(http.StatusBadRequest, &w, err) {
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
		err = errors.New(strings.Join(errMsgs, ", "))
		if utils.IsNotNilAndWriteErrResp(http.StatusBadRequest, &w, err) {
			return
		}
	}

	key, err := utils.Produce[Request]("input-harga", reqBody, utils.ValueByteConverterFunc[Request](ValueByteConvertion))
	if utils.IsNotNilAndWriteErrResp(http.StatusInternalServerError, &w, err) {
		return
	}

	response := model.ResponseTmpl{IsError: false, ReffID: key}
	responseJSON, err := json.Marshal(response)
	if utils.IsNotNilAndWriteErrResp(http.StatusInternalServerError, &w, err) {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
