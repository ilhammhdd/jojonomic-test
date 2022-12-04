package main

import (
	"net/http"

	"github.com/ilhammhdd/jojonomic_test/model"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

func HandleCekSalo(w http.ResponseWriter, r *http.Request) {
	reqBody, err := utils.ReadRequestBodyJSON[requestBody](r)
	if utils.IsNotNilAndWriteErrResp(http.StatusBadRequest, &w, err) {
		return
	}

	rekening, err := SelectRekening(reqBody.Norek)
	if err != nil && err == ErrNoAccountExists {
		respBodyJSON, err1 := model.NewResponseTmplJSON(false, "", err, nil)
		if err1 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(respBodyJSON)
		return
	} else if utils.IsNotNilAndWriteErrResp(http.StatusInternalServerError, &w, err) {
		return
	}

	respBodyJSON, err := BuildResponseBody(rekening)
	if utils.IsNotNilAndWriteErrResp(http.StatusInternalServerError, &w, err) {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respBodyJSON)
}
