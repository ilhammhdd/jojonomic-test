package main

import (
	"net/http"
	"time"

	"github.com/ilhammhdd/jojonomic_test/model"
	"github.com/ilhammhdd/jojonomic_test/utils"
)

func HandleCekMutasi(w http.ResponseWriter, r *http.Request) {
	reqBody, err := utils.ReadRequestBodyJSON[inRequestBody](r)
	if utils.IsNotNilAndWriteErrResp(http.StatusBadRequest, &w, err) {
		return
	}

	transactions, err := SelectTransactions(time.Unix(reqBody.StartDate, 0).UTC(), time.Unix(reqBody.EndDate, 0).UTC())
	if err != nil && err == ErrNoTransaksi {
		respBodyJSON, err1 := model.NewResponseTmplJSON(false, "", err, nil)
		if err1 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(respBodyJSON)
	}

	respBodyJSON, err := model.NewResponseTmplJSON(false, "", nil, transactions)
	if utils.IsNotNilAndWriteErrResp(http.StatusInternalServerError, &w, err) {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respBodyJSON)
}
