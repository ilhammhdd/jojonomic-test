package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/ilhammhdd/jojonomic_test/model"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

func HandleTopup(w http.ResponseWriter, r *http.Request) {
	reqBody, err := utils.ReadRequestBodyJSON[model.TransaksiRequest](r)
	if utils.IsNotNilAndWriteErrResp(http.StatusBadRequest, &w, err) {
		return
	}

	err = validateReqBody(reqBody)
	if utils.IsNotNilAndWriteErrResp(http.StatusBadRequest, &w, err) {
		return
	}

	key, err := Topup(reqBody)
	if err != nil {
		if err == ErrPriceDoesntMatch {
			utils.IsNotNilAndWriteErrResp(http.StatusBadRequest, &w, err)
		} else {
			utils.IsNotNilAndWriteErrResp(http.StatusInternalServerError, &w, err)
		}
		return
	}

	resp := model.ResponseTmpl{IsError: false, ReffID: key}
	respJSON, err := json.Marshal(resp)
	if utils.IsNotNilAndWriteErrResp(http.StatusInternalServerError, &w, err) {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)
}

func validateReqBody(reqBody *model.TransaksiRequest) error {
	var errMsgs []string
	if !utils.CheckMaxDecimalPlaces(3, reqBody.Gram) {
		errMsgs = append(errMsgs, "gram minimal and multiplication of 0.001")
	}
	if reqBody.Gram <= 0 {
		errMsgs = append(errMsgs, "gram must be > 0")
	}
	if reqBody.NoRek == "" {
		errMsgs = append(errMsgs, "norek can't be empty")
	}

	if len(errMsgs) > 0 {
		return errors.New(strings.Join(errMsgs, ", "))
	}
	return nil
}
